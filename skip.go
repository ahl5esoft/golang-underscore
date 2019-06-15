package underscore

import "reflect"

func (m enumerable) Skip(count int) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			iterator := m.GetEnumerator()
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					for ; count > 0; count-- {
						if !iterator.MoveNext() {
							return
						}
					}

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
