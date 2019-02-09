package underscore

import (
	"reflect"
)

// First is 获取第一个元素
func First(source, match interface{}) {
	rv := reflect.ValueOf(match)
	if rv.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

	length, getKeyValue := parseSource(source)
	if length == 0 {
		return
	}

	valueRV, _ := getKeyValue(0)
	rv.Elem().Set(valueRV)
}

func (m *query) First() IQuery {
	First(m.Source, &m.Source)
	return m
}
