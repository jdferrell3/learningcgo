package main

import (
	"fmt"
)

func FindCommon(a []int, b []int) ([]int) {
	common := []int{}
	seen := make(map[int]bool)

	for _, item := range a {
		seen[item] = true
	}

	for _, item := range b {
		if _, ok := seen[item]; ok {
			common = append(common, item)
		}
	}

	return common
}

func main() {
	fmt.Print(FindCommon([]int{1, 2, 3}, []int{3, 4, 5}))
}
