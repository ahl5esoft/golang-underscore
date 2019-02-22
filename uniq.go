package underscore

import (
	"reflect"
)

// Uniq is 去重
func Uniq(source, selector interface{}) interface{} {
	if selector == nil {
		selector = func(value, _ interface{}) facade {
			return facade{reflect.ValueOf(value)}
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
	getPropertyRV := PropertyRV(property)
	return Uniq(source, func(value, _ interface{}) facade {
		return facade{
			getPropertyRV(value),
		}
	})
}

func (m *query) Uniq(selector interface{}) IQuery {
	m.Source = Uniq(m.Source, selector)
	return m
}

func (m *query) UniqBy(property string) IQuery {
	m.Source = UniqBy(m.Source, property)
	return m
}
