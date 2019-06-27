package underscore

import "reflect"

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
