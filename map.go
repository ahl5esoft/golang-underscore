package underscore

import (
	"errors"
	"reflect"
)

func Map(source, selector interface{}) (interface{}, error) {
	selectorRV := reflect.ValueOf(selector)
	if selectorRV.Kind() != reflect.Func {
		return nil, errors.New("underscore: Map's selector is not func")
	}

	arrRV := makeSliceRVWithElem(selectorRV.Type().Out(0), 0)
	err := each(source, func (args []reflect.Value) (bool, reflect.Value) {
		values := selectorRV.Call(args)
		if !isErrorRVValid(values[1]) {
			arrRV = reflect.Append(arrRV, values[0])
		}

		return false, values[1]
	})
	if err == nil && arrRV.IsValid() {
		return arrRV.Interface(), nil
	}

	return nil, err
}

//chain
func (this *Query) Map(selector interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = Map(this.source, selector)
	}
	return this
}