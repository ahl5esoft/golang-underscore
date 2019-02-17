package underscore

import (
	"reflect"
)

// Index is 转化为indexSelector筛选出的值为key的map
func Index(source, indexSelector, result interface{}) {
	resultRV := reflect.ValueOf(result)
	if resultRV.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

	var tempRV reflect.Value
	each(source, indexSelector, func(indexRV, valueRV, _ reflect.Value) bool {
		if !tempRV.IsValid() {
			tempRT := reflect.MapOf(indexRV.Type(), valueRV.Type())
			tempRV = reflect.MakeMap(tempRT)
		}

		tempRV.SetMapIndex(indexRV, valueRV)
		return false
	})
	if tempRV.IsValid() {
		resultRV.Elem().Set(tempRV)
	}
}

// IndexBy is 转化为property值的map
func IndexBy(source interface{}, property string, res interface{}) {
	getPropertyRV := PropertyRV(property)
	Index(source, func(value, _ interface{}) facade {
		return facade{
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
