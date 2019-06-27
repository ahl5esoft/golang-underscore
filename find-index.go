package underscore

import "reflect"

func (m enumerable) FindIndex(predicate interface{}) int {
	iterator := m.GetEnumerator()
	predicateRV := reflect.ValueOf(predicate)
	index := 0
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		returnRVs := predicateRV.Call([]reflect.Value{
			iterator.GetValue(),
			iterator.GetKey(),
		})
		if returnRVs[0].Bool() {
			return index
		}

		index++
	}

	return -1
}

func (m enumerable) FindIndexBy(dict map[string]interface{}) int {
	return m.FindIndex(func(v, _ interface{}) bool {
		return IsMatch(v, dict)
	})
}
