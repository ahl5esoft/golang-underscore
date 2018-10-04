package underscore

import (
	"reflect"
)

// Map 映射
func Map(source, selector interface{}) interface{} {
	var arrRV reflect.Value
	each(source, selector, func(resRV, valueRV, _ reflect.Value) bool {
		if !arrRV.IsValid() {
			arrRT := reflect.SliceOf(resRV.Type())
			arrRV = reflect.MakeSlice(arrRT, 0, 0)
		}

		arrRV = reflect.Append(arrRV, resRV)
		return false
	})
	if arrRV.IsValid() {
		return arrRV.Interface()
	}

	return nil
}

// Map is IQuery's method
func (q *Query) Map(selector interface{}) IQuery {
	q.source = Map(q.source, selector)
	return q
}
