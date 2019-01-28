package underscore

import (
	"reflect"
)

// All of the values in the `source` pass the `predicate` truth test
// @source		map or array
// @predicate	func(value or item, key or index) bool
func All(source, predicate interface{}) bool {
	var ok bool
	each(source, predicate, func(resRV, _, _ reflect.Value) bool {
		ok = resRV.Bool()
		return !ok
	})
	return ok
}

// AllBy will stop traversing the `source` if a false element is found
// @source		map or array
// @properties	map[string]interface{}
func AllBy(source interface{}, properties map[string]interface{}) bool {
	return All(source, func(value, _ interface{}) bool {
		return IsMatch(value, properties)
	})
}

func (m *query) All(predicate interface{}) bool {
	return All(m.Source, predicate)
}

func (m *query) AllBy(properties map[string]interface{}) bool {
	return AllBy(m.Source, properties)
}
