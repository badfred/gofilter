gofilter
========

gofilter creates a filter function that remove elements from a slice where a provided predicate returns false. The difference to common filters is that this one can handle arbitrary typed slices by using reflect.MakeFunc from go1.1. A generic map function is also available

Requires go1.1. Contains some code duplication in SetFilter and SetMap I sometime should remove.

Example
-------
```Go
package main

import (
	"github.com/badfred/gofilter"
	"fmt"
	"log"
)

func intToString(a int) string {
	return fmt.Sprint(a)
}

func largerThan5(a int) bool {
	return a > 5
}


func main() {
	
	input := []int{3, 4, 5, 6, 7}
	
	var intFilter func([]int, func(int) bool) []int
	err := filter.SetFilter(&intFilter)
	if err != nil {
		log.Fatalf("filter.SetFilter: %v", err)
	}
	
	var intMap func([]int, func(int) string) []string
	err = filter.SetMap(&intMap)
	if err != nil {
		log.Fatalf("filter.SetMap: %v", err)
	}
	
	
	fmt.Printf("intFilter(%#v, largerThan5) :  %#v\n", input, intFilter(input, largerThan5))
	fmt.Printf("intMap(%#v, intToString)    :  %#v\n", input, intMap(input, intToString))
	
}
```

Copyright
---------

Copyright: MIT license
