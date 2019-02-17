package underscore

import (
	"reflect"
)

// Values is 字典的所有value
func Values(source, values interface{}) {
	sourceRV := reflect.ValueOf(source)
	if sourceRV.Kind() != reflect.Map {
		return
	}

	Map(source, func(value, _ interface{}) facade {
		return facade{reflect.ValueOf(value)}
	}, values)
}

func (m *query) Values() IQuery {
	Values(m.Source, &m.Source)
	return m
}
