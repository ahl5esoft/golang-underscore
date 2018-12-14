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

	origin := Clone(memo)
	iteratorRV := reflect.ValueOf(iterator)
	memoRV := reflect.ValueOf(memo)
	for i := 0; i < length; i++ {
		valueRV, keyRV := getKeyValue(i)
		returnRVs := iteratorRV.Call(
			[]reflect.Value{
				memoRV,
				valueRV,
				keyRV,
			},
		)
		memoRV = returnRVs[0]
	}

	if memoRV.IsValid() {
		return memoRV.Interface()
	}

	return origin
}

// Reduce is IQuery's method
func (q *Query) Reduce(iterator, memo interface{}) IQuery {
	q.source = Reduce(q.source, iterator, memo)
	return q
}
