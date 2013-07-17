package gofilter

import (
	"errors"
	"reflect"
)

/*
SetFilter stores a function where filter points to. When filter is called afterwards,
it filters a slice by applying a predicate on each element. The returned slice will
consist of the elements where the predicate returned true.
filter must be a pointer to a function of type func([]A, func(A) bool) []A,
e.g. func([]int, func(int) bool) []int. SetFilter will use the type information
to create a matching filter function.
*/
func SetFilter(filter interface{}) error {

	pointerValue := reflect.ValueOf(filter)
	if pointerValue.Kind() != reflect.Ptr {
		return errors.New("filter must be a pointer")
	}

	filterValue := pointerValue.Elem()
	if filterValue.Kind() != reflect.Func {
		return errors.New("filter must be a pointer to a function")
	}

	filterType := filterValue.Type()

	if filterType.NumIn() != 2 {
		return errors.New("filter must have exactly two input parameters")
	}

	if filterType.NumOut() != 1 {
		return errors.New("filter must have exactly one output parameter")
	}

	outputType := filterType.Out(0)
	if outputType.Kind() != reflect.Slice {
		return errors.New("filter must return a slice")
	}

	arrayType := filterType.In(0)
	if arrayType.Kind() != reflect.Slice {
		return errors.New("first argument of filter must be a slice")
	}

	predicateType := filterType.In(1)
	if predicateType.Kind() != reflect.Func {
		return errors.New("second argument of filter must be a function")
	}

	if predicateType.NumIn() != 1 {
		return errors.New("predicate must have exactly one input parameter")
	}

	if predicateType.NumOut() != 1 {
		return errors.New("predicate must have exactly one output parameter")
	}

	if predicateType.Out(0).Kind() != reflect.Bool {
		return errors.New("predicate must return a Boolean value")
	}

	inputArrayElType := arrayType.Elem()
	inputPredicateArgType := predicateType.In(0)
	outputArrayElType := outputType.Elem()

	if inputArrayElType.ConvertibleTo(inputPredicateArgType) == false {
		return errors.New("input array elements cannot be converted to predicate parameter")
	}

	if inputArrayElType.ConvertibleTo(outputArrayElType) == false {
		return errors.New("input array elemenents cannot be converted to output array elements")
	}

	filterImpl := func(in []reflect.Value) []reflect.Value {
		a := in[0]
		predicate := in[1]

		result := reflect.MakeSlice(outputType, 0, 0)

		for i := 0; i < a.Len(); i++ {
			el := a.Index(i)
			predicateArg := []reflect.Value{el.Convert(inputPredicateArgType)}
			predicateResult := predicate.Call(predicateArg)
			if predicateResult[0].Bool() {
				result = reflect.Append(result, el.Convert(outputArrayElType))
			}
		}

		return []reflect.Value{result}
	}

	function := reflect.MakeFunc(filterType, filterImpl)
	filterValue.Set(function)

	return nil
}
