package gofilter

import (
	"testing"
)

func isEven(i int) bool {
	return i%2 == 0
}

func isGreaterThan5(f float32) bool {
	return f > 5
}

func TestSetIntFilter(t *testing.T) {

	var isEvenTest = struct {
		in  []int
		out []int
	}{
		[]int{1, 2, 3, 4, 5},
		[]int{2, 4},
	}

	var intFilter func([]int, func(int) bool) []int
	err := SetFilter(&intFilter)
	if err != nil {
		t.Fatalf("gofilter.SetFilter: %v", err)
	}

	got := intFilter(isEvenTest.in, isEven)
	if len(got) != len(isEvenTest.out) {
		t.Fatalf("expected %#v, got %#v\n", isEvenTest.out, got)
	}
	for i, _ := range isEvenTest.out {
		if got[i] != isEvenTest.out[i] {
			t.Fatalf("expected %#v, got %#v\n", isEvenTest.out, got)
		}
	}

}

func TestSetFloat32Filter(t *testing.T) {

	var isGreaterThan5Test = struct {
		in  []float32
		out []float32
	}{[]float32{4, 4.5, 5, 5.5, 6, 6.5, 7}, []float32{5.5, 6, 6.5, 7}}

	var float32Filter func([]float32, func(float32) bool) []float32
	err := SetFilter(&float32Filter)
	if err != nil {
		t.Fatalf("gofilter.SetFilter: %v", err)
	}

	got := float32Filter(isGreaterThan5Test.in, isGreaterThan5)
	if len(got) != len(isGreaterThan5Test.out) {
		t.Fatalf("expected %#v, got %#v\n", isGreaterThan5Test.out, got)
	}
	for i, _ := range isGreaterThan5Test.out {
		if got[i] != isGreaterThan5Test.out[i] {
			t.Fatalf("expected %#v, got %#v\n", isGreaterThan5Test.out, got)
		}
	}

}
