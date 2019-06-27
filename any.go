package underscore

import "reflect"

func (m enumerable) Any(predicate interface{}) bool {
	iterator := m.GetEnumerator()
	predicateRV := reflect.ValueOf(predicate)
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		returnRVs := predicateRV.Call([]reflect.Value{
			iterator.GetValue(),
			iterator.GetKey(),
		})
		if returnRVs[0].Bool() {
			return true
		}
	}

	return false
}

func (m enumerable) AnyBy(dict map[string]interface{}) bool {
	return m.Any(func(v, _ interface{}) bool {
		return IsMatch(v, dict)
	})
}
