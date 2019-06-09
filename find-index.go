package underscore

import "reflect"

func (m *query) FindIndex(predicate interface{}) int {
	index := -1

	if !IsArray(m.Source) {
		return index
	}

	each(m.Source, predicate, func(okRV, _, keyRV reflect.Value) bool {
		ok := okRV.Bool()
		if ok {
			index = int(keyRV.Int())
		}
		return ok
	})

	return index
}

func (m *query) FindIndexBy(properties map[string]interface{}) int {
	return m.FindIndex(func(item interface{}, _ int) bool {
		return IsMatch(item, properties)
	})
}

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
