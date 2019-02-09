package underscore

import "reflect"

// Last is 最后元素
func Last(source, last interface{}) {
	lastRV := reflect.ValueOf(last)
	if lastRV.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

	length, getKeyValue := parseSource(source)
	if length == 0 {
		return
	}

	valueRV, _ := getKeyValue(length - 1)
	lastRV.Elem().Set(valueRV)
}

func (m *query) Last() IQuery {
	Last(m.Source, &m.Source)
	return m
}
