package underscore

import (
	"reflect"
)

// Values is 字典的所有value
func Values(source interface{}) interface{} {
	sourceRV := reflect.ValueOf(source)
	if sourceRV.Kind() != reflect.Map {
		return nil
	}

	return Map(source, func(value, _ interface{}) facade {
		return facade{reflect.ValueOf(value)}
	})
}

func (m *query) Values() IQuery {
	m.Source = Values(m.Source)
	return m
}
