package underscore

import "reflect"

func getRealRV(v interface{}) reflect.Value {
	value := reflect.ValueOf(v)
	if value.Type() == valueType {
		value = v.(reflect.Value)
	}

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Type() == facadeType {
		value = value.Interface().(facade).Real
	}

	return value
}

func getFuncReturnRV(selectorValue reflect.Value, enumerator IEnumerator) reflect.Value {
	return getRealRV(
		selectorValue.Call(
			[]reflect.Value{
				enumerator.GetValue(),
				enumerator.GetKey(),
			},
		)[0],
	)
}
