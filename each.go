package underscore

import "reflect"

func (m enumerable) Each(action interface{}) {
	iterator := m.GetEnumerator()
	actionValue := reflect.ValueOf(action)
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		actionValue.Call(
			[]reflect.Value{
				iterator.GetValue(),
				iterator.GetKey(),
			},
		)
	}
}
