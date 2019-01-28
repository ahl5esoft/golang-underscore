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
	var first interface{}
	First(source, &first)
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

func (m *query) Object() IQuery {
	if m.IsParallel {
		m.Source = objectAsParallel(m.Source)
	} else {
		m.Source = Object(m.Source)
	}

	return m
}
