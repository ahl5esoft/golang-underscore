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

// Index is IQuery's method
func (q *Query) Index(indexSelector interface{}) IQuery {
	q.source = Index(q.source, indexSelector)
	return q
}

// IndexBy is IQuery's method
func (q *Query) IndexBy(property string) IQuery {
	q.source = IndexBy(q.source, property)
	return q
}
