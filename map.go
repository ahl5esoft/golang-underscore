package underscore

import (
	"reflect"
)

// Map 映射
func Map(source, selector, result interface{}) {
	rv := reflect.ValueOf(result)
	if rv.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

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
		rv.Elem().Set(arrRV)
	}
}

// MapBy is 从source中取出所有property
func MapBy(source interface{}, property string, result interface{}) {
	getPropertyRV := PropertyRV(property)
	Map(source, func(value, _ interface{}) Facade {
		return Facade{
			getPropertyRV(value),
		}
	}, result)
}

func (m *query) Map(selector interface{}) IQuery {
	Map(m.Source, selector, &m.Source)
	return m
}

func (m *query) MapBy(property string) IQuery {
	MapBy(m.Source, property, &m.Source)
	return m
}
