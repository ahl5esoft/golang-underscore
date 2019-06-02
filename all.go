package underscore

import (
	"reflect"
)

func (m *query) All(predicate interface{}) bool {
	var ok bool
	each(m.Source, predicate, func(resRV, _, _ reflect.Value) bool {
		ok = resRV.Bool()
		return !ok
	})
	return ok
}

func (m *query) AllBy(properties map[string]interface{}) bool {
	return m.All(func(value, _ interface{}) bool {
		return IsMatch(value, properties)
	})
}

func (m enumerable) All(predicate interface{}) bool {
	iterator := m.GetEnumerator()
	predicateRV := reflect.ValueOf(predicate)
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		returnRVs := predicateRV.Call([]reflect.Value{
			iterator.GetValue(),
			iterator.GetKey(),
		})
		if !returnRVs[0].Bool() {
			return false
		}
	}

	return true
}

func (m enumerable) AllBy(dict map[string]interface{}) bool {
	return m.All(func(v, _ interface{}) bool {
		return IsMatch(v, dict)
	})
}
