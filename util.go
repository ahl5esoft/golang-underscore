package underscore

import (
	"errors"
	"reflect"
)

func getFieldValue(entity interface{}, fieldName string) (interface{}, error) {
	rv := reflect.ValueOf(entity).FieldByName(fieldName)
	if rv.IsValid() {
		return rv.Interface(), nil
	}

	return nil, errors.New("invalid field: [" + fieldName + "]")
}