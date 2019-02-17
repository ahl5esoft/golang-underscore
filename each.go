package underscore

import (
	"reflect"
	"sync"
)

var facadeRT = reflect.TypeOf(facade{})

// Each will iterate over a list of elements
func Each(source, iterator interface{}) {
	each(source, iterator, nil)
}

func (m *query) Each(iterator interface{}) IQuery {
	if m.IsParallel {
		eachAsParallel(m.Source, iterator)
	} else {
		each(m.Source, iterator, nil)
	}

	return m
}

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
			if returnRVs[0].Type() == facadeRT {
				returnRVs[0] = returnRVs[0].Interface().(facade).Real
			}

			if predicate(returnRVs[0], valueRV, keyRV) {
				break
			}
		}
	}
}

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
		case reflect.Array, reflect.Slice:
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
