package underscore

import (
	"fmt"
	"reflect"
)

// Find is 根据断言获取元素
func Find(source, predicate, matcher interface{}) {
	rv := reflect.ValueOf(matcher)
	if rv.Kind() != reflect.Ptr {
		panic(
			fmt.Sprintf("receive type must be a pointer: `Find(_, _, macther interface{})`"),
		)
	}

	var ok bool
	each(source, predicate, func(resRV, valueRV, _ reflect.Value) bool {
		ok = resRV.Bool()
		if ok {
			rv.Elem().Set(valueRV)
		}
		return ok
	})
}

// FindBy is 根据属性值获取元素
func FindBy(source interface{}, properties map[string]interface{}, matcher interface{}) {
	Find(source, func(value, _ interface{}) bool {
		return IsMatch(value, properties)
	}, matcher)
}

func (m *query) Find(predicate interface{}) IQuery {
	Find(m.Source, predicate, &m.Source)
	return m
}

func (m *query) FindBy(properties map[string]interface{}) IQuery {
	FindBy(m.Source, properties, &m.Source)
	return m
}
