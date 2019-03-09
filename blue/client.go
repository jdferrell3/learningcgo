package main

import (
	"expvar"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

const MaxDelay = 5
const MinDelay = 1
const TotalRequests = 1000

func myClient(i int, wg *sync.WaitGroup, url string) {
	fmt.Printf("client %d started\n", i)
	client := &http.Client{
		Timeout: time.Second * 35,
	}

	uri := fmt.Sprintf("%s/%d", url, i)

	// Get the data
	// resp, err := client.Get(uri)
	// if nil != err {
	// 	fmt.Print(err)
	// } else {
	// 	fmt.Printf("%d\n", resp.StatusCode)
	// 	results <- resp.StatusCode
	// }

	req, err := http.NewRequest("GET", uri, nil)
	if nil != err {
		log.Fatal(err)
	}

	// Don't re-use the request
	req.Close = true

	resp, err := client.Do(req)
	if nil != err {
		fmt.Print(err)
	} else {
		fmt.Printf("client %d got back %d\n", i, resp.StatusCode)
		results <- resp.StatusCode
	}
	defer resp.Body.Close()

	fmt.Printf("client %d ended\n", i)
	wg.Done()
}

var results = make(chan int, 10)

func main() {
	// expvar uses /debug/vars as the default URI
	// get the expvar handler and change the URI
	h := expvar.Handler()
	http.HandleFunc("/status", h.ServeHTTP)
	go http.ListenAndServe(":8080", http.DefaultServeMux)
	counts := expvar.NewMap("counters")

	// Test server to emulate some of the server-side requirements
	codes := []int{200, 404, 500}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%+v\n", r.URL)

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		// random response code
		n := r1.Int() % len(codes)

		// random int between min and max
		t := rand.Intn(MaxDelay-MinDelay) + MinDelay

		fmt.Printf("Delaying response for %d\n", t)
		time.Sleep(time.Duration(t) * time.Second)
		w.WriteHeader(codes[n])
		w.Write([]byte(`hello`))
	}))
	defer ts.Close()

	var wg sync.WaitGroup

	// counter for requests, each request is made by a go routine
	j := 1

	// start 10 clients
	for j <= 10 {
		wg.Add(1)
		go myClient(j, &wg, ts.URL)
		time.Sleep(10 * time.Millisecond)
		j++
	}

	// start another client when we get a result from the channel
	for j <= TotalRequests {
		select {
		case r := <-results:
			// update our expvar counts map
			counts.Add(fmt.Sprintf("%d", r), int64(1))

			// we got a result, start a new client
			wg.Add(1)
			go myClient(j, &wg, ts.URL)
			j++
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
	wg.Wait()

	fmt.Println("All go routines finished executing")
}
