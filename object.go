package underscore

import (
	"reflect"
	"sync"
)

// Object 将二维数组转化为字典
func Object(source interface{}) interface{} {
	var mapRv reflect.Value
	each(source, func(item, _ interface{}) {
		rv := reflect.ValueOf(item)
		keyRv := rv.Index(0).Elem()
		valueRv := rv.Index(1).Elem()
		if !mapRv.IsValid() {
			mapRv = reflect.MakeMap(
				reflect.MapOf(
					keyRv.Type(),
					valueRv.Type(),
				),
			)
		}

		mapRv.SetMapIndex(keyRv, valueRv)
	}, nil)

	if mapRv.IsValid() {
		return mapRv.Interface()
	}

	return nil
}

func objectAsParallel(source interface{}) interface{} {
	first := First(source)
	if first == nil || Size(first) != 2 {
		return nil
	}

	firstRv := reflect.ValueOf(first)
	keyRv := firstRv.Index(0)
	valueRv := firstRv.Index(1)
	mapRv := reflect.MakeMap(
		reflect.MapOf(
			keyRv.Type(),
			valueRv.Type(),
		),
	)

	var mutex sync.RWMutex
	eachAsParallel(source, func(item, _ interface{}) {
		rv := reflect.ValueOf(item)
		mutex.Lock()
		defer mutex.Unlock()

		mapRv.SetMapIndex(
			rv.Index(0),
			rv.Index(1),
		)
	})

	return mapRv.Interface()
}

// Object is Queryer'e method
func (m *Query) Object() Queryer {
	if m.isParallel {
		m.source = objectAsParallel(m.source)
	} else {
		m.source = Object(m.source)
	}

	return m
}
