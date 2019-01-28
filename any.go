package underscore

import (
	"reflect"
)

// Any is if any of the values in the `source` pass the `predicate` truth test
// @source		map or array
// @predicate	func(value or item, key or index) bool
func Any(source, predicate interface{}) bool {
	var ok bool
	each(source, predicate, func(resRV, _, _ reflect.Value) bool {
		ok = resRV.Bool()
		return ok
	})
	return ok
}

// AnyBy will stop traversing the `source` if a true element is found
// @source		map or array
func AnyBy(source interface{}, properties map[string]interface{}) bool {
	return Any(source, func(value, _ interface{}) bool {
		return IsMatch(value, properties)
	})
}

func (m *query) Any(predicate interface{}) bool {
	return Any(m.Source, predicate)
}

func (m *query) AnyBy(properties map[string]interface{}) bool {
	return AnyBy(m.Source, properties)
}
