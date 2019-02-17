package underscore

import (
	"reflect"
)

// Map 映射
func Map(source, selector, result interface{}) {
	resultRV := reflect.ValueOf(result)
	if resultRV.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

	var tempRV reflect.Value
	each(source, selector, func(resRV, valueRV, _ reflect.Value) bool {
		if !tempRV.IsValid() {
			tempRT := reflect.SliceOf(resRV.Type())
			tempRV = reflect.MakeSlice(tempRT, 0, 0)
		}

		tempRV = reflect.Append(tempRV, resRV)
		return false
	})
	if tempRV.IsValid() {
		resultRV.Elem().Set(tempRV)
	}
}

// MapBy is 从source中取出所有property
func MapBy(source interface{}, property string, result interface{}) {
	getPropertyRV := PropertyRV(property)
	Map(source, func(value, _ interface{}) facade {
		return facade{
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
