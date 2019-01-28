package underscore

import (
	"reflect"
)

// Index is 转化为indexSelector筛选出的值为key的map
func Index(source, indexSelector interface{}) interface{} {
	var dictRV reflect.Value
	each(source, indexSelector, func(indexRV, valueRV, _ reflect.Value) bool {
		if !dictRV.IsValid() {
			dictRT := reflect.MapOf(indexRV.Type(), valueRV.Type())
			dictRV = reflect.MakeMap(dictRT)
		}

		dictRV.SetMapIndex(indexRV, valueRV)
		return false
	})
	if dictRV.IsValid() {
		return dictRV.Interface()
	}

	return nil
}

// IndexBy is 转化为property值的map
func IndexBy(source interface{}, property string) interface{} {
	getPropertyRV := PropertyRV(property)
	return Index(source, func(value, _ interface{}) Facade {
		return Facade{
			getPropertyRV(value),
		}
	})
}

func (m *query) Index(indexSelector interface{}) IQuery {
	m.Source = Index(m.Source, indexSelector)
	return m
}

func (m *query) IndexBy(property string) IQuery {
	m.Source = IndexBy(m.Source, property)
	return m
}
