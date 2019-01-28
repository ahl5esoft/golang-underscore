package underscore

import (
	"fmt"
	"reflect"
)

// First is 获取第一个元素
func First(source, matcher interface{}) {
	rv := reflect.ValueOf(matcher)
	if rv.Kind() != reflect.Ptr {
		panic(
			fmt.Sprintf("receive type must be a pointer: `First(_, macther interface{})`"),
		)
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
