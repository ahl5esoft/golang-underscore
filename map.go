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

// MapBy 从source中取出所有property
func MapBy(source interface{}, property string) interface{} {
	getPropertyRV := PropertyRV(property)
	return Map(source, func(value, _ interface{}) Facade {
		rv, _ := getPropertyRV(value)
		return Facade{rv}
	})
}

// Map is Queryer's method
func (q *Query) Map(selector interface{}) Queryer {
	q.source = Map(q.source, selector)
	return q
}

// MapBy is Queryer's method
func (q *Query) MapBy(property string) Queryer {
	q.source = MapBy(q.source, property)
	return q
}
