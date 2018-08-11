package underscore

import (
	"reflect"
)

// Object is 转为对象
func Object(source interface{}) interface{} {
	var mapRv reflect.Value
	each(source, func(value, _ interface{}) {
		rv := reflect.ValueOf(value)
		if !mapRv.IsValid() {
			mapRv = reflect.MakeMap(
				reflect.MapOf(
					rv.Index(0).Elem().Type(),
					rv.Index(1).Elem().Type(),
				),
			)
		}

		mapRv.SetMapIndex(
			rv.Index(0).Elem(),
			rv.Index(1).Elem(),
		)
	}, nil)

	if mapRv.IsValid() {
		return mapRv.Interface()
	}

	return nil
}
