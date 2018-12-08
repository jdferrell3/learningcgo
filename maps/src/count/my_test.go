package main

import (
	"testing"
)

type M map[string]int

type TStruct struct {
	String string
	Map    M
}

func TestCountChars(t *testing.T) {
	var tstructs = []TStruct{
		TStruct{
			String: "hello",
			Map: M{string('h'): 1,
				string('l'): 2},
		},
		TStruct{
			String: "gggooo",
			Map: M{string('g'): 3,
				string('o'): 3},
		}}

	for _, tstruct := range tstructs {
		chars := countChars(tstruct.String)
		for k, v := range tstruct.Map {
			if chars[k] != v {
				t.Fatal("unexpected count")
			}
		}
	}
}
