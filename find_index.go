package underscore

import (
	"errors"
	"reflect"
)

func FindIndex(source, predicate interface{}) (int, error) {
	index := -1

	predicateRV := reflect.ValueOf(predicate)
	if predicateRV.Kind() != reflect.Func {
		return index, errors.New("underscore: FindIndex's predicate is not func")
	}

	sourceRV := reflect.ValueOf(source)
	if !(sourceRV.Kind() == reflect.Array || sourceRV.Kind() == reflect.Slice) {
		return index, nil
	}

	each(source, func (args []reflect.Value) (bool, reflect.Value) {
		values := predicateRV.Call(args)
		ok := values[0].Bool()
		if ok {
			index = int(args[1].Int())
		}
		return ok, values[1]
	})
	return index, nil
}