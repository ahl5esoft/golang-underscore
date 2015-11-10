package underscore

import (
	"errors"
	"reflect"
)

func getPropertyValue(entity interface{}, property string) (interface{}, error) {
	rv := reflect.ValueOf(entity).FieldByName(property)
	if rv.IsValid() {
		return rv.Interface(), nil
	}

	return nil, errors.New("invalid field: [" + property + "]")
}

func makeArray(rt reflect.Type) reflect.Value {
	return reflect.MakeSlice(rt, 0, 0)
}