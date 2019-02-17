package underscore

import (
	"reflect"
)

// First is 获取第一个元素
func First(source, first interface{}) {
	firstRV := reflect.ValueOf(first)
	if firstRV.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

	length, getKeyValue := parseSource(source)
	if length == 0 {
		return
	}

	valueRV, _ := getKeyValue(0)
	firstRV.Elem().Set(valueRV)
}

func (m *query) First() IQuery {
	First(m.Source, &m.Source)
	return m
}
