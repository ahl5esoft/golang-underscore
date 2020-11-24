package underscore

import "reflect"

func (m enumerable) Select(selector interface{}) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			iterator := m.GetEnumerator()
			selectorRV := reflect.ValueOf(selector)
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					if ok = iterator.MoveNext(); ok {
						keyRV = iterator.GetKey()
						valueRV = getReturnValue(selectorRV, iterator)
					}

					return
				},
			}
		},
	}
}

func (m enumerable) SelectBy(fieldName string) IEnumerable {
	getter := FieldValue(fieldName)
	return m.Select(func(value, _ interface{}) facade {
		return facade{
			getter(value),
		}
	})
}
