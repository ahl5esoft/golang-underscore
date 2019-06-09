package underscore

import "reflect"

func (m enumerable) Select(selector interface{}) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			index := 0
			iterator := m.GetEnumerator()
			selectorRV := reflect.ValueOf(selector)
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					if ok = iterator.MoveNext(); ok {
						valueRV = selectorRV.Call([]reflect.Value{
							iterator.GetValue(),
							iterator.GetKey(),
						})[0]
						keyRV = reflect.ValueOf(index)
						index++
					}

					return
				},
			}
		},
	}
}

func (m enumerable) SelectBy(fieldName string) IEnumerable {
	getter := PropertyRV(fieldName)
	return m.Select(func(value, _ interface{}) facade {
		return facade{
			getter(value),
		}
	})
}
