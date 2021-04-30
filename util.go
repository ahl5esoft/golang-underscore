package underscore

import "reflect"

func getRealValue(v interface{}) reflect.Value {
	value := reflect.ValueOf(v)
	if value.Type() == valueType {
		value = v.(reflect.Value)
	}

	if value.Type() == facadeType {
		value = value.Interface().(facade).Real
	}

	return value
}

func getReturnValue(selectorValue reflect.Value, enumerator IEnumerator) reflect.Value {
	return getRealValue(
		selectorValue.Call(
			[]reflect.Value{
				enumerator.GetValue(),
				enumerator.GetKey(),
			},
		)[0],
	)
}
