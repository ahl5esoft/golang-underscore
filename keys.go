package underscore

import (
	"reflect"
)

// Keys is 获取map的所有key
func Keys(source interface{}) interface{} {
	sourceRV := reflect.ValueOf(source)
	if sourceRV.Kind() != reflect.Map {
		return nil
	}

	return Map(source, func(_, key interface{}) Facade {
		return Facade{reflect.ValueOf(key)}
	})
}

func (m *query) Keys() IQuery {
	m.Source = Keys(m.Source)
	return m
}
