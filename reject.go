package underscore

import (
	"reflect"
)

// Reject is 排除
func Reject(source, predicate, result interface{}) {
	rv := reflect.ValueOf(result)
	if rv.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

	rv.Elem().Set(
		reflect.ValueOf(
			filter(source, predicate, false),
		),
	)
}

// RejectBy is 根据属性排除
func RejectBy(source interface{}, properties map[string]interface{}, result interface{}) {
	Reject(source, func(value, _ interface{}) bool {
		return IsMatch(value, properties)
	}, result)
}

func (m *query) Reject(predicate interface{}) IQuery {
	Reject(m.Source, predicate, &m.Source)
	return m
}

func (m *query) RejectBy(properties map[string]interface{}) IQuery {
	RejectBy(m.Source, properties, &m.Source)
	return m
}
