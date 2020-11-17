package underscore

import "reflect"

func (m enumerable) Aggregate(fn interface{}, memo interface{}) IEnumerable {
	fnValue := reflect.ValueOf(fn)
	iterator := m.GetEnumerator()
	memoValue := reflect.ValueOf(memo)
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		memoValue = fnValue.Call(
			[]reflect.Value{
				memoValue,
				iterator.GetValue(),
				iterator.GetKey(),
			},
		)[0]
	}
	return chainFromValue(memoValue)
}
