package underscore

import "reflect"

func (m enumerable) Find(predicate interface{}) IEnumerable {
	iterator := m.GetEnumerator()
	predicateRV := reflect.ValueOf(predicate)
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		returnRVs := predicateRV.Call([]reflect.Value{
			iterator.GetValue(),
			iterator.GetKey(),
		})
		if returnRVs[0].Bool() {
			return chainFromValue(
				iterator.GetValue(),
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
