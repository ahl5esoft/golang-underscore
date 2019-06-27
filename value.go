package underscore

import "reflect"

func (m enumerable) Value(res interface{}) {
	resRV := reflect.ValueOf(res)
	switch resRV.Elem().Kind() {
	case reflect.Array, reflect.Slice:
		m.valueToArrayOrSlice(resRV)
	case reflect.Map:
		m.valueToMap(resRV)
	default:
		if nullIterator, ok := m.GetEnumerator().(nullEnumerator); ok {
			if rv := nullIterator.GetValue(); rv.IsValid() {
				resRV.Elem().Set(rv)
			}
		}
	}
}

func (m enumerable) valueToArrayOrSlice(resRV reflect.Value) {
	iterator := m.GetEnumerator()
	sliceRV := resRV.Elem()
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		sliceRV = reflect.Append(
			sliceRV,
			iterator.GetValue(),
		)
	}

	resRV.Elem().Set(sliceRV)
}

func (m enumerable) valueToMap(resRV reflect.Value) {
	iterator := m.GetEnumerator()
	mapRV := resRV.Elem()
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		mapRV.SetMapIndex(
			iterator.GetKey(),
			iterator.GetValue(),
		)
	}
	resRV.Elem().Set(mapRV)
}
