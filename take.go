package underscore

import "reflect"

func (m enumerable) Take(count int) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			iterator := m.GetEnumerator()
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					if count <= 0 {
						return
					}

					count--
					if ok = iterator.MoveNext(); ok {
						valueRV = iterator.GetValue()
						keyRV = iterator.GetKey()
					}

					return
				},
			}
		},
	}
}
