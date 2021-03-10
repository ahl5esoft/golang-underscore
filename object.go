package underscore

import (
	"reflect"
)

func (m enumerable) Object() IEnumerable {
	iterator := m.GetEnumerator()
	return enumerable{
		Enumerator: func() IEnumerator {
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					if ok = iterator.MoveNext(); ok {
						keyRV = iterator.GetValue().Index(0)
						valueRV = iterator.GetValue().Index(1)
						if keyRV.Kind() == reflect.Interface {
							keyRV = keyRV.Elem()
							valueRV = valueRV.Elem()
						}
					}

					return
				},
			}
		},
	}
}
