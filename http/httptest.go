package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

const PinnedCertHash = "60b87575447dcba2a36b7d11ac09fb24a9db406fee12d2cc90180517616e8a18"
const TestURL = "https://httpbin.org/base64/SFRUUEJJTiBpcyBhd2Vzb21l"
const TestURLResp = "HTTPBIN is awesome"

var (
	ErrBadCert = errors.New("cert does not match pinned fingerprint")
)

func main() {
	// _, err := httpsClient(TestURL)
	// if nil != err {
	// 	log.Println(err)
	// }

	// _, err = httpsClientPinned(TestURL, PinnedCertHash)
	// if nil != err {
	// 	log.Println(err)
	// }

	// _, err = httpsClientPinned(TestURL, "foo")
	// if nil == err {
	// 	log.Println("This should have erred")
	// }

	data, err := PostJSON("https://httpbin.org/post")
	if nil != err {
		log.Println(err)
	}
	log.Println(string(data))

	data, err = PostJSONViaRequest("https://httpbin.org/post", []byte(`{"string":"there"}`))
	if nil != err {
		log.Println(err)
	}
	log.Println(string(data))
	var f interface{}
	err = json.Unmarshal(data, &f)
	fmt.Printf("%v", f.json)
}

func PostJSONViaRequest(url string, sendData []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(sendData))
	if nil != err {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 10}

	client.Transport = &http.Transport{
		TLSHandshakeTimeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	return data, nil
}

func PostJSON(url string) ([]byte, error) {
	// curl -X POST "https://httpbin.org/post" -H "Content-Type: application/json, accept: application/json" --data  '{"string":"hello"}'
	client := &http.Client{
		Timeout: time.Second * 10}

	client.Transport = &http.Transport{
		TLSHandshakeTimeout: 5 * time.Second,
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer([]byte(`{"string":"hello"}`)))
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	return data, nil
}

func Downloader() {
	tmpfile, err := ioutil.TempFile("", "example")
	if nil != err {
		log.Fatal(err)
	}

	if err := tmpfile.Close(); nil != err {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	uri := "https://httpbin.org/bytes/500"
	err = DownloadFile(tmpfile.Name(), uri)
	if nil != err {
		panic(err)
	}
}

func httpsClientPinned(url string, pinnedCert string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10}

	client.Transport = &http.Transport{
		DialTLS:             makePinnedDialer(pinnedCert, false),
		TLSHandshakeTimeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	return data, nil
}

func httpsClient(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10}

	client.Transport = &http.Transport{
		TLSHandshakeTimeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	return data, nil
}

func DownloadFile(filepath string, url string) error {
	log.Println(filepath)
	// Create the file
	out, err := os.Create(filepath)
	if nil != err {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if nil != err {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if nil != err {
		return err
	}

	return nil
}

type Dialer func(network string, addr string) (net.Conn, error)

func makePinnedDialer(pinnedCertHash string, skipCAVerification bool) Dialer {
	return func(network string, addr string) (net.Conn, error) {
		c, err := tls.Dial(network, addr, &tls.Config{InsecureSkipVerify: skipCAVerification})
		if nil != err {
			return nil, err
		}

		connstate := c.ConnectionState()
		for _, peercert := range connstate.PeerCertificates {
			// RawSubjectPublicKeyInfo is the DER encoded SubjectPublicKeyInfo
			// we'll hash the public key, this makes cert renewals easier to handle
			// than hashing the entire cert
			hashBytes := sha256.Sum256(peercert.RawSubjectPublicKeyInfo)
			hash := hex.EncodeToString(hashBytes[:])

			// log.Printf("%+v\n", peercert.Subject)
			// log.Printf("%+v\n", peercert.RawSubjectPublicKeyInfo)
			// log.Printf("sha256 hash of peercert.RawSubjectPublicKeyInfo: %#v\n", hash)

			if strings.Compare(hash, pinnedCertHash) == 0 {
				// log.Printf("Pinned cert found")
				return c, nil
			}
		} // end PeerCertificates

		return nil, ErrBadCert
	} // end func()
} // end makePinnedDialer()
