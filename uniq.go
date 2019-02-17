package underscore

import (
	"reflect"
)

// Uniq is 去重
func Uniq(source, selector, result interface{}) {
	resultRV := reflect.ValueOf(result)
	if resultRV.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

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
		resultRV.Elem().Set(arrRV)
	}
}

// UniqBy is 根据某个属性去重
func UniqBy(source interface{}, property string, result interface{}) {
	getPropertyRV := PropertyRV(property)
	Uniq(source, func(value, _ interface{}) facade {
		return facade{
			getPropertyRV(value),
		}
	}, result)
}

func (m *query) Uniq(selector interface{}) IQuery {
	Uniq(m.Source, selector, &m.Source)
	return m
}

func (m *query) UniqBy(property string) IQuery {
	UniqBy(m.Source, property, &m.Source)
	return m
}
