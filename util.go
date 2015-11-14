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

func makeSliceRV(rt reflect.Type) reflect.Value {
	return reflect.MakeSlice(rt, 0, 0)
}

func makeSliceRVWithElem(elemRT reflect.Type) reflect.Value {
	sliceType := reflect.SliceOf(elemRT)
	return makeSliceRV(sliceType)
}

func makeMapRV(keyRT, valueRT reflect.Type) reflect.Value {
	mapRT := reflect.MapOf(keyRT, valueRT)
	return reflect.MakeMap(mapRT)
}

func makeGroupRV(keyRT, elemRT reflect.Type) reflect.Value {
	sliceRT := reflect.SliceOf(elemRT)
	return makeMapRV(keyRT, sliceRT)
}