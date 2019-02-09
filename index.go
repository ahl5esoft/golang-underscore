package underscore

import (
	"reflect"
)

// Index is 转化为indexSelector筛选出的值为key的map
func Index(source, indexSelector, res interface{}) {
	resRV := reflect.ValueOf(res)
	if resRV.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

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
		resRV.Elem().Set(dictRV)
	}
}

// IndexBy is 转化为property值的map
func IndexBy(source interface{}, property string, res interface{}) {
	getPropertyRV := PropertyRV(property)
	Index(source, func(value, _ interface{}) Facade {
		return Facade{
			getPropertyRV(value),
		}
	}, res)
}

func (m *query) Index(indexSelector interface{}) IQuery {
	Index(m.Source, indexSelector, &m.Source)
	return m
}

func (m *query) IndexBy(property string) IQuery {
	IndexBy(m.Source, property, &m.Source)
	return m
}
