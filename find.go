package underscore

import (
	"reflect"
)

// Find is 根据断言获取元素
func Find(source, predicate interface{}) interface{} {
	var ok bool
	var v interface{}
	each(source, predicate, func(resRV, valueRV, _ reflect.Value) bool {
		ok = resRV.Bool()
		if ok {
			v = valueRV.Interface()
		}
		return ok
	})
	return v
}

// FindBy is 根据属性值获取元素
func FindBy(source interface{}, properties map[string]interface{}) interface{} {
	return Find(source, func(value, _ interface{}) bool {
		return IsMatch(value, properties)
	})
}

func (m *query) Find(predicate interface{}) IQuery {
	m.Source = Find(m.Source, predicate)
	return m
}

func (m *query) FindBy(properties map[string]interface{}) IQuery {
	m.Source = FindBy(m.Source, properties)
	return m
}
