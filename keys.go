package underscore

import (
	"reflect"
)

// Keys is 获取map的所有key
func Keys(source, keys interface{}) {
	sourceRV := reflect.ValueOf(source)
	if sourceRV.Kind() != reflect.Map {
		return
	}

	Map(source, func(_, key interface{}) Facade {
		return Facade{reflect.ValueOf(key)}
	}, keys)
}

func (m *query) Keys() IQuery {
	Keys(m.Source, &m.Source)
	return m
}
