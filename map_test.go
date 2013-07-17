package gofilter

import (
	"fmt"
	"testing"
)

func timesTwo(i int) int {
	return 2 * i
}

func intToString(i int) string {
	return fmt.Sprint(i)
}

func TestSetIntMap(t *testing.T) {

	var timesTwoTest = struct {
		in  []int
		out []int
	}{
		[]int{1, 2, 3, 4, 5},
		[]int{2, 4, 6, 8, 10},
	}

	var intMap func([]int, func(int) int) []int
	err := SetMap(&intMap)
	if err != nil {
		t.Fatalf("filter.SetMap: %v", err)
	}

	got := intMap(timesTwoTest.in, timesTwo)
	if len(got) != len(timesTwoTest.out) {
		t.Fatalf("expected %#v, got %#v\n", timesTwoTest.out, got)
	}
	for i, _ := range timesTwoTest.out {
		if got[i] != timesTwoTest.out[i] {
			t.Fatalf("expected %#v, got %#v\n", timesTwoTest.out, got)
		}
	}

}

func TestMapIntToString(t *testing.T) {

	var mapIntToStringTest = struct {
		in  []int
		out []string
	}{
		[]int{4, 5, 6, 7},
		[]string{"4", "5", "6", "7"},
	}

	var intToStringMap func([]int, func(int) string) []string
	err := SetMap(&intToStringMap)
	if err != nil {
		t.Fatalf("filter.SetMap: %v", err)
	}

	got := intToStringMap(mapIntToStringTest.in, intToString)
	if len(got) != len(mapIntToStringTest.out) {
		t.Fatalf("expected %#v, got %#v\n", mapIntToStringTest.out, got)
	}
	for i, _ := range mapIntToStringTest.out {
		if got[i] != mapIntToStringTest.out[i] {
			t.Fatalf("expected %#v, got %#v\n", mapIntToStringTest.out, got)
		}
	}

}
