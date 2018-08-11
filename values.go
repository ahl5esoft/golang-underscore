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

	return Map(source, func(value, _ interface{}) Facade {
		return Facade{reflect.ValueOf(value)}
	})
}

// Values is Queryer's method
func (q *Query) Values() Queryer {
	q.source = Values(q.source)
	return q
}
