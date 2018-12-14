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

// Any is IQuery's method
func (q *Query) Any(predicate interface{}) IQuery {
	q.source = Any(q.source, predicate)
	return q
}

// AnyBy is IQuery's method
func (q *Query) AnyBy(properties map[string]interface{}) IQuery {
	q.source = AnyBy(q.source, properties)
	return q
}
