package underscore

import "reflect"

func (m enumerable) Keys() IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			index := 0
			iterator := m.GetEnumerator()
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					if ok = iterator.MoveNext(); ok {
						valueRV = iterator.GetKey()
						keyRV = reflect.ValueOf(index)
						index++
					}

					return
				},
			}
		},
	}
}
