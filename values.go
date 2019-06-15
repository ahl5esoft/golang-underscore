package underscore

import (
	"reflect"
)

func (m *query) Values() IQuery {
	sourceRV := reflect.ValueOf(m.Source)
	if sourceRV.Kind() == reflect.Map {
		m.Source = m.Map(func(value, _ interface{}) facade {
			return facade{reflect.ValueOf(value)}
		})
	} else {
		m.Source = nil
	}

	return m
}

func (m enumerable) Values() IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			index := 0
			iterator := m.GetEnumerator()
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					if ok = iterator.MoveNext(); ok {
						valueRV = iterator.GetValue()
						keyRV = reflect.ValueOf(index)
						index++
					}

					return
				},
			}
		},
	}
}
