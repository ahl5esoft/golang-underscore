package underscore

import "reflect"

func (m enumerable) SelectMany(selector interface{}) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			var tempIterator IEnumerator
			iterator := m.GetEnumerator()
			selectorRV := reflect.ValueOf(selector)
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					for !ok {
						if tempIterator == nil {
							ok = iterator.MoveNext()
							if !ok {
								return
							}

							selectorResultRV := getReturnValue(selectorRV, iterator)
							tempIterator = chainFromValue(selectorResultRV).GetEnumerator()
						}

						if ok = tempIterator.MoveNext(); ok {
							keyRV = tempIterator.GetKey()
							valueRV = tempIterator.GetValue()
						} else {
							tempIterator = nil
						}
					}
					return
				},
			}
		},
	}
}

func (m enumerable) SelectManyBy(fieldName string) IEnumerable {
	getter := FieldValue(fieldName)
	return m.SelectMany(func(value, _ interface{}) facade {
		return facade{
			getter(value),
		}
	})
}
