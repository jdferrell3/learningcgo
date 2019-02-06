package main

import (
	"reflect"
	"testing"
)

func TestCountChars(t *testing.T) {
	tstructs := []struct {
		expected []int
		listA    []int
		listB    []int
	}{
		{
			[]int{1, 4, 9},
			[]int{1, 3, 4, 6, 7, 9},
			[]int{1, 2, 4, 5, 9, 10},
		},
		{
			[]int{1, 2, 9, 10, 12},
			[]int{1, 2, 9, 10, 11, 12},
			[]int{0, 1, 2, 3, 4, 5, 8, 9, 10, 12, 14, 15},
		},
		{
			[]int{},
			[]int{0, 1, 2, 3, 4, 5},
			[]int{6, 7, 8, 9, 10, 11},
		},
	}

	for _, tt := range tstructs {
		actual := FindCommon(tt.listA, tt.listB)
		if !reflect.DeepEqual(tt.expected, actual) {
			t.Fatalf("expected: %+v; actual: %+v", tt.expected, actual)
		}
	}
}
