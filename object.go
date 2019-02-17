package underscore

import (
	"reflect"
	"sync"
)

// Object 将二维数组转化为字典
func Object(source, result interface{}) {
	resultRV := reflect.ValueOf(result)
	if resultRV.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

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
		resultRV.Elem().Set(tempRV)
	}
}

func (m *query) Object() IQuery {
	if m.IsParallel {
		objectAsParallel(m.Source, &m.Source)
	} else {
		Object(m.Source, &m.Source)
	}

	return m
}

func objectAsParallel(source, result interface{}) {
	resultRV := reflect.ValueOf(result)
	if resultRV.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

	var first interface{}
	First(source, &first)
	if first == nil || Size(first) != 2 {
		return
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

	resultRV.Elem().Set(tempRV)
}
