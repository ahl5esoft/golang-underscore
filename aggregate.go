package underscore

import "reflect"

func (m enumerable) Aggregate(memo interface{}, fn interface{}) IEnumerable {
	fnRV := reflect.ValueOf(fn)
	iterator := m.GetEnumerator()
	memoRV := reflect.ValueOf(memo)
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		memoRV = fnRV.Call([]reflect.Value{
			memoRV,
			iterator.GetValue(),
			iterator.GetKey(),
		})[0]
	}
	return chainFromRV(memoRV)
}
