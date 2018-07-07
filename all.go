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

// All is Queryer's method
func (q *Query) All(predicate interface{}) Queryer {
	q.source = All(q.source, predicate)
	return q
}

// AllBy is Queryer's method
func (q *Query) AllBy(properties map[string]interface{}) Queryer {
	q.source = AllBy(q.source, properties)
	return q
}
