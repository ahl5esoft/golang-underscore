package underscore

import "reflect"

func (m *query) Where(predicate interface{}) IQuery {
	m.Source = filter(m.Source, predicate, true)
	return m
}

func (m *query) WhereBy(properties map[string]interface{}) IQuery {
	m.Source = m.Where(func(value, _ interface{}) bool {
		return IsMatch(value, properties)
	})
	return m
}

func (m enumerable) Where(predicate interface{}) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			iterator := m.GetEnumerator()
			predicateRV := reflect.ValueOf(predicate)
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					for ok = iterator.MoveNext(); ok; ok = iterator.MoveNext() {
						valueRV = iterator.GetValue()
						keyRV = iterator.GetKey()
						if predicateRV.Call([]reflect.Value{valueRV, keyRV})[0].Bool() {
							return
						}
					}

					return
				},
			}
		},
	}
}

func (m enumerable) WhereBy(dict map[string]interface{}) IEnumerable {
	return m.Where(func(v, _ interface{}) bool {
		return IsMatch(v, dict)
	})
}
