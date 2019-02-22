package underscore

import (
	"reflect"
	"sync"
)

// Object 将二维数组转化为字典
func Object(source interface{}) interface{} {
	var tempRV reflect.Value
	each(source, func(item, _ interface{}) {
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

	if tempRV.IsValid() {
		return tempRV.Interface()
	}
	return nil
}

func (m *query) Object() IQuery {
	if m.IsParallel {
		m.Source = objectAsParallel(m.Source)
	} else {
		m.Source = Object(m.Source)
	}

	return m
}

func objectAsParallel(source interface{}) interface{} {
	first := First(source)
	if first == nil || Size(first) != 2 {
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
