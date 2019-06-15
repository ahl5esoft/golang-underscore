package underscore

import (
	"reflect"
)

// Reduce is 聚合
func Reduce(source, iterator, memo interface{}) interface{} {
	length, getKeyValue := parseSource(source)
	if length == 0 {
		return memo
	}

	iteratorRV := reflect.ValueOf(iterator)
	memoRV := reflect.ValueOf(memo)
	for i := 0; i < length; i++ {
		valueRV, keyRV := getKeyValue(i)
		memoRV = iteratorRV.Call(
			[]reflect.Value{
				memoRV,
				valueRV,
				keyRV,
			},
		)[0]
	}

	return memoRV.Interface()
}

func (m *query) Reduce(iterator, memo interface{}) IQuery {
	m.Source = Reduce(m.Source, iterator, memo)
	return m
}

func (m enumerable) Reduce(memo interface{}, fn interface{}) IEnumerable {
	return m.Aggregate(memo, fn)
}
