package underscore

import (
	"errors"
	"reflect"
)

func each(source interface{}, iterator func([]reflect.Value) (bool, reflect.Value)) error {
	if source == nil {
		return nil
	}

	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			if sourceRV.Len() == 0 {
				return nil
			}

			for i := 0; i < sourceRV.Len(); i++ {
				isBreak, errRV := iterator(
					[]reflect.Value{
						sourceRV.Index(i),
						reflect.ValueOf(i),
					},
				)
				
				if isErrorRVValid(errRV) {
					return errRV.Interface().(error)
				}

				if isBreak {
					return nil
				}
			}
		case reflect.Map:
			keyRVs := sourceRV.MapKeys()
			if len(keyRVs) == 0 {
				return nil
			}

			for i := 0; i < len(keyRVs); i++ {
				isBreak, errRV := iterator(
					[]reflect.Value{
						sourceRV.MapIndex(keyRVs[i]),
						keyRVs[i],
					},
				)
				
				if isErrorRVValid(errRV) {
					return errRV.Interface().(error)
				}

				if isBreak {
					return nil
				}
			}
	}
	return nil
}

func getPropertyRV(entityRV reflect.Value, property string) (reflect.Value, error) {
	var err error
	rv := entityRV.FieldByName(property)
	if !rv.IsValid() {
		err = errors.New("invalid field: [" + property + "]")
	}

	return rv, err
}

func getPropertyValue(entity interface{}, property string) (interface{}, error) {
	rv, err := getPropertyRV(
		reflect.ValueOf(entity),
		property,
	)
	if err == nil {
		return rv.Interface(), nil
	}
	
	return nil, err
}

func isErrorRVValid(errRV reflect.Value) bool {
	return errRV.IsValid() && !errRV.IsNil()
}

func makeSliceRV(rt reflect.Type, length int) reflect.Value {
	return reflect.MakeSlice(rt, length, length)
}

func makeSliceRVWithElem(elemRT reflect.Type, length int) reflect.Value {
	sliceType := reflect.SliceOf(elemRT)
	return makeSliceRV(sliceType, length)
}

func makeMapRV(keyRT, valueRT reflect.Type) reflect.Value {
	mapRT := reflect.MapOf(keyRT, valueRT)
	return reflect.MakeMap(mapRT)
}

func makeGroupRV(keyRT, elemRT reflect.Type) reflect.Value {
	sliceRT := reflect.SliceOf(elemRT)
	return makeMapRV(keyRT, sliceRT)
}