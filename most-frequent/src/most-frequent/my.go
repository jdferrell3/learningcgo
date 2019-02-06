package main

import "fmt"

func main() {}

func MostFreq(l []int) int {
	numbers := make(map[int]int)
	for _, n := range l {
		// fmt.Println(n)
		numbers[n]++
	}

	most := 0
	mostCount := 0
	for num, count := range numbers {
		fmt.Println(num, count)
		if count > mostCount {
			most = num
			mostCount = count
		}
	}

	return most
}