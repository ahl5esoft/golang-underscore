package underscore

import (
	"reflect"
)

// Map 映射
func Map(source, selector interface{}) interface{} {
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
		return tempRV.Interface()
	}

	return nil
}

// MapBy is 从source中取出所有property
func MapBy(source interface{}, property string) interface{} {
	getPropertyRV := PropertyRV(property)
	return Map(source, func(value, _ interface{}) facade {
		return facade{
			getPropertyRV(value),
		}
	})
}

func (m *query) Map(selector interface{}) IQuery {
	m.Source = Map(m.Source, selector)
	return m
}

func (m *query) MapBy(property string) IQuery {
	m.Source = MapBy(m.Source, property)
	return m
}

func (m enumerable) Map(selector interface{}) IEnumerable {
	return m.Select(selector)
}

func (m enumerable) MapBy(fieldName string) IEnumerable {
	return m.SelectBy(fieldName)
}
