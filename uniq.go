package underscore

import (
	"reflect"
)

// Uniq is 去重
func Uniq(source, selector interface{}) interface{} {
	if selector == nil {
		selector = func(value, _ interface{}) Facade {
			return Facade{reflect.ValueOf(value)}
		}
	}

	var mapRV reflect.Value
	var arrRV reflect.Value
	each(source, selector, func(resRV, valueRv, _ reflect.Value) bool {
		if !mapRV.IsValid() {
			mapRT := reflect.MapOf(resRV.Type(), reflect.TypeOf(false))
			mapRV = reflect.MakeMap(mapRT)

			arrRT := reflect.SliceOf(valueRv.Type())
			arrRV = reflect.MakeSlice(arrRT, 0, 0)
		}

		mapValueRV := mapRV.MapIndex(resRV)
		if !mapValueRV.IsValid() {
			mapRV.SetMapIndex(resRV, reflect.ValueOf(true))
			arrRV = reflect.Append(arrRV, valueRv)
		}
		return false
	})
	if mapRV.IsValid() {
		return arrRV.Interface()
	}

	return nil
}

// UniqBy is 根据某个属性去重
func UniqBy(source interface{}, property string) interface{} {
	getProeprtyRV := PropertyRV(property)
	return Uniq(source, func(value, _ interface{}) Facade {
		rv, _ := getProeprtyRV(value)
		return Facade{rv}
	})
}

// Uniq is Queryer's method
func (q *Query) Uniq(selector interface{}) Queryer {
	q.source = Uniq(q.source, selector)
	return q
}

// UniqBy is Queryer's method
func (q *Query) UniqBy(property string) Queryer {
	q.source = UniqBy(q.source, property)
	return q
}
