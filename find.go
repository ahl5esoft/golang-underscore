package underscore

import (
	"reflect"
)

func (m *query) Find(predicate interface{}) IQuery {
	var ok bool
	each(m.Source, predicate, func(resRV, valueRV, _ reflect.Value) bool {
		ok = resRV.Bool()
		if ok {
			m.Source = valueRV.Interface()
		}
		return ok
	})
	return m
}

func (m *query) FindBy(properties map[string]interface{}) IQuery {
	return m.Find(func(value, _ interface{}) bool {
		return IsMatch(value, properties)
	})
}

func (m enumerable) Find(predicate interface{}) IEnumerable {
	iterator := m.GetEnumerator()
	predicateRV := reflect.ValueOf(predicate)
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		returnRVs := predicateRV.Call([]reflect.Value{
			iterator.GetValue(),
			iterator.GetKey(),
		})
		if returnRVs[0].Bool() {
			return Chain2(
				iterator.GetValue().Interface(),
			)
		}
	}

	return nilEnumerable
}

func (m enumerable) FindBy(dict map[string]interface{}) IEnumerable {
	return m.Find(func(v, _ interface{}) bool {
		return IsMatch(v, dict)
	})
}
