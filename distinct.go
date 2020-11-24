package underscore

import "reflect"

func (m enumerable) Distinct(selector interface{}) IEnumerable {
	if selector == nil {
		selector = func(value, _ interface{}) facade {
			return facade{
				reflect.ValueOf(value),
			}
		}
	}

	return enumerable{
		Enumerator: func() IEnumerator {
			iterator := m.GetEnumerator()
			selectValue := reflect.ValueOf(selector)
			set := make(map[interface{}]bool)
			return &enumerator{
				MoveNextFunc: func() (valueValue reflect.Value, keyValue reflect.Value, ok bool) {
					for ok = iterator.MoveNext(); ok; ok = iterator.MoveNext() {
						valueValue = iterator.GetValue()
						keyValue = iterator.GetKey()
						v := getReturnValue(selectValue, iterator).Interface()
						if _, has := set[v]; !has {
							set[v] = true
							return
						}
					}
					return
				},
			}
		},
	}
}

func (m enumerable) DistinctBy(fieldName string) IEnumerable {
	getter := FieldValue(fieldName)
	return m.Distinct(func(value, _ interface{}) facade {
		return facade{
			getter(value),
		}
	})
}
