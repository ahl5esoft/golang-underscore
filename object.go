package underscore

import (
	"reflect"
	"sync"
)

func (m *query) Object() IQuery {
	if m.IsParallel {
		m.Source = objectAsParallel(m.Source)
	} else {
		var tempRV reflect.Value
		each(m.Source, func(item, _ interface{}) {
			rv := reflect.ValueOf(item)
			keyRv := rv.Index(0).Elem()
			valueRv := rv.Index(1).Elem()
			if !tempRV.IsValid() {
				tempRV = reflect.MakeMap(
					reflect.MapOf(
						keyRv.Type(),
						valueRv.Type(),
					),
				)
			}

			tempRV.SetMapIndex(keyRv, valueRv)
		}, nil)
		m.Source = tempRV.Interface()
	}

	return m
}

func objectAsParallel(source interface{}) interface{} {
	first := First(source)
	if first == nil || Chain(first).Size() != 2 {
		return nil
	}

	firstRv := reflect.ValueOf(first)
	tempRV := reflect.MakeMap(
		reflect.MapOf(
			firstRv.Index(0).Type(),
			firstRv.Index(1).Type(),
		),
	)

	var mutex sync.RWMutex
	eachAsParallel(source, func(item, _ interface{}) {
		rv := reflect.ValueOf(item)

		mutex.Lock()
		defer mutex.Unlock()

		tempRV.SetMapIndex(
			rv.Index(0),
			rv.Index(1),
		)
	})

	return tempRV.Interface()
}

func (m enumerable) Object() IEnumerable {
	iterator := m.GetEnumerator()
	return enumerable{
		Enumerator: func() IEnumerator {
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					if ok = iterator.MoveNext(); ok {
						keyRV = iterator.GetValue().Index(0).Elem()
						valueRV = iterator.GetValue().Index(1).Elem()
					}

					return
				},
			}
		},
	}
}
