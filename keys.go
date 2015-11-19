package underscore

import (
	"errors"
	"reflect"
)

func Keys(source interface{}) (interface{}, error) {
	return mapFromEach(source, 1)
}

func mapFromEach(source interface{}, index int) (interface{}, error) {
	rv := reflect.ValueOf(source)
	if rv.Kind() != reflect.Map {
		return nil, errors.New("underscore: Keys's source is not map")
	}

	var arrRV reflect.Value
	err := each(source, func (args []reflect.Value) (bool, reflect.Value) {
		if !arrRV.IsValid() {
			arrRV = makeSliceRVWithElem(args[index].Type(), 0)
		}

		arrRV = reflect.Append(arrRV, args[index])
		return false, reflect.ValueOf(nil)
	})
	if err == nil && arrRV.IsValid() {
		return arrRV.Interface(), nil
	}

	return nil, err
}

//Chain
func (this *Query) Keys() Queryer {
	if this.err == nil {
		this.source, this.err = Keys(this.source)
	}
	return this
}