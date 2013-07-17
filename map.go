package gofilter

import (
	"errors"
	"reflect"
)

/*
SetMap stores a function where m points to. When m is called afterwards,
it maps a slice by applying a mapping function f on each element. The returned slice will
consist of the same number of elements, but with f applied on them.
m must be a pointer to a function of type func([]A, func(A) B) []B,
e.g. func([]int, func(int) string) []string. SetMap will use the type information
to create the matching map function. If the type signature is not compatible an
error will be returned.
*/
func SetMap(m interface{}) error {

	pointerValue := reflect.ValueOf(m)
	if pointerValue.Kind() != reflect.Ptr {
		return errors.New("map must be a pointer")
	}

	mapValue := pointerValue.Elem()
	if mapValue.Kind() != reflect.Func {
		return errors.New("map must be a pointer to a function")
	}

	mapType := mapValue.Type()

	if mapType.NumIn() != 2 {
		return errors.New("map must have exactly two input parameters")
	}

	if mapType.NumOut() != 1 {
		return errors.New("map must have exactly one output parameter")
	}

	outputType := mapType.Out(0)
	if outputType.Kind() != reflect.Slice {
		return errors.New("map must return a slice")
	}

	arrayType := mapType.In(0)
	if arrayType.Kind() != reflect.Slice {
		return errors.New("first argument of map must be a slice")
	}

	fType := mapType.In(1)
	if fType.Kind() != reflect.Func {
		return errors.New("second argument of map must be a function")
	}

	if fType.NumIn() != 1 {
		return errors.New("f must have exactly one input parameter")
	}

	if fType.NumOut() != 1 {
		return errors.New("f must have exactly one output parameter")
	}

	inputArrayElType := arrayType.Elem()
	inputFArgType := fType.In(0)
	oututFType := fType.Out(0)
	outputArrayElType := outputType.Elem()

	if inputArrayElType.ConvertibleTo(inputFArgType) == false {
		return errors.New("input array elements cannot be converted to f's parameter")
	}

	if oututFType.ConvertibleTo(outputArrayElType) == false {
		return errors.New("f result type cannot be converted to output array elements")
	}

	mapImpl := func(in []reflect.Value) []reflect.Value {
		a := in[0]
		f := in[1]

		result := reflect.MakeSlice(outputType, 0, a.Len())

		for i := 0; i < a.Len(); i++ {
			el := a.Index(i)
			fArg := []reflect.Value{el.Convert(inputFArgType)}
			fResult := f.Call(fArg)[0]
			result = reflect.Append(result, fResult.Convert(outputArrayElType))
		}

		return []reflect.Value{result}
	}

	function := reflect.MakeFunc(mapType, mapImpl)
	mapValue.Set(function)

	return nil
}
