package underscore

import "reflect"

func getRealRV(v interface{}) reflect.Value {
	rv := reflect.ValueOf(v)
	if rv.Type() == rtOfRV {
		rv = v.(reflect.Value)
	}

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Type() == facadeRT {
		rv = rv.Interface().(facade).Real
	}

	return rv
}

func getFuncReturnRV(selectorRV reflect.Value, enumerator IEnumerator) reflect.Value {
	return getRealRV(
		selectorRV.Call([]reflect.Value{
			enumerator.GetValue(),
			enumerator.GetKey(),
		})[0],
	)
}
