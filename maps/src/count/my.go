package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countChars(s string) map[string]int {
	chars := make(map[string]int)
	for _, c := range s {
		// fmt.Println(c)
		chars[string(c)]++
	}
	return chars
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		chars := countChars(s)

		fmt.Println(s)
		for k, v := range chars {
			fmt.Printf("'%s': %d, ", k, v)
		}
		fmt.Printf("\n\n")
	}
}
