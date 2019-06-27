package underscore

import "reflect"

func (m enumerable) Each(action interface{}) {
	iterator := m.GetEnumerator()
	actionRV := reflect.ValueOf(action)
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		actionRV.Call([]reflect.Value{
			iterator.GetValue(),
			iterator.GetKey(),
		})
	}
}
