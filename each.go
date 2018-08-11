package underscore

import (
	"reflect"
	"sync"
)

// Each will iterate over a list of elements
// @source		map or array
// @iterator	func(value or item, key or index) bool
func Each(source, iterator interface{}) {
	each(source, iterator, nil)
}

// each will iterate over a list of elements
// @source		map or array
// @iterator	func(value or item, key or index) bool
// @predicate	stop traversing if pass the `predicate` truth test
func each(source interface{}, iterator interface{}, predicate func(reflect.Value, reflect.Value, reflect.Value) bool) {
	length, getKeyValue := parseSource(source)
	if length == 0 {
		return
	}

	if predicate == nil {
		predicate = func(resRV, _, _ reflect.Value) bool {
			if resRV.Kind() == reflect.Bool {
				return resRV.Bool()
			}

			return false
		}
	}

	iteratorRV := reflect.ValueOf(iterator)
	for i := 0; i < length; i++ {
		valueRV, keyRV := getKeyValue(i)
		returnRVs := iteratorRV.Call(
			[]reflect.Value{valueRV, keyRV},
		)
		if len(returnRVs) > 0 {
			resRV := returnRVs[0]
			if resRV.Type() == FacadeRt {
				resRV = resRV.Interface().(Facade).Real
			}

			if predicate(resRV, valueRV, keyRV) {
				break
			}
		}
	}
}

// eachAsParallel will parallel iterate over a list of elements
// @source		map or array
// @iterator	func(value or item, key or index) bool
func eachAsParallel(source interface{}, iterator interface{}) {
	length, getKeyValue := parseSource(source)
	if length == 0 {
		return
	}

	var task sync.WaitGroup
	task.Add(length)

	iteratorRV := reflect.ValueOf(iterator)
	for i := 0; i < length; i++ {
		go func(index int) {
			valueRV, keyRV := getKeyValue(index)
			iteratorRV.Call(
				[]reflect.Value{valueRV, keyRV},
			)

			task.Done()
		}(i)
	}

	task.Wait()
}

func parseSource(source interface{}) (int, func(i int) (reflect.Value, reflect.Value)) {
	if source != nil {
		sourceRV := reflect.ValueOf(source)
		switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			return sourceRV.Len(), func(i int) (reflect.Value, reflect.Value) {
				return sourceRV.Index(i), reflect.ValueOf(i)
			}
		case reflect.Map:
			keyRVs := sourceRV.MapKeys()
			return len(keyRVs), func(i int) (reflect.Value, reflect.Value) {
				return sourceRV.MapIndex(keyRVs[i]), keyRVs[i]
			}
		}
	}
	return 0, nil
}

// Each is Queryer's method
func (q *Query) Each(iterator interface{}) Queryer {
	if q.isParallel {
		eachAsParallel(q.source, iterator)
	} else {
		each(q.source, iterator, nil)
	}

	return q
}
