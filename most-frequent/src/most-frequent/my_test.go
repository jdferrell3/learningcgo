package main

import (
	"testing"
)

func TestMostFreq(t *testing.T) {
	tstructs := []struct {
		expected int
		list    []int
	}{
		{
			1,
			[]int{1, 3, 1, 3, 2, 1},
		},
		{
			3,
			[]int{3, 3, 1, 3, 2, 1},
		},
		{
			0,
			[]int{0},
 		},
 		{
 			-1,
 			[]int{0, -1, 10, 10, -1, 10, -1, -1, -1, 1},
		},
	}

	for _, tt := range tstructs {
		actual := MostFreq(tt.list)
		if actual != tt.expected {
			t.Fatalf("fail - expected: %d; actual: %d\n", tt.expected, actual)
		}
	}
}
