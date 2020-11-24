package underscore

import "reflect"

func (m enumerable) Index(keySelector interface{}) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			iterator := m.GetEnumerator()
			keySelectorRV := reflect.ValueOf(keySelector)
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					if ok = iterator.MoveNext(); ok {
						keyRV = getReturnValue(keySelectorRV, iterator)
						valueRV = iterator.GetValue()
					}

					return
				},
			}
		},
	}
}

func (m enumerable) IndexBy(fieldName string) IEnumerable {
	getter := FieldValue(fieldName)
	return m.Index(func(value, _ interface{}) facade {
		return facade{
			getter(value),
		}
	})
}
