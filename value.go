package underscore

import "reflect"

func (m enumerable) Value(res interface{}) {
	resValue := reflect.ValueOf(res)
	switch resValue.Elem().Kind() {
	case reflect.Array, reflect.Slice:
		m.valueToArrayOrSlice(resValue)
	case reflect.Map:
		m.valueToMap(resValue)
	default:
		if nullIterator, ok := m.GetEnumerator().(nullEnumerator); ok {
			if value := nullIterator.GetValue(); value.IsValid() {
				resValue.Elem().Set(value)
			}
		}
	}
}

func (m enumerable) valueToArrayOrSlice(resValue reflect.Value) {
	iterator := m.GetEnumerator()
	sliceValue := resValue.Elem()
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		sliceValue = reflect.Append(
			sliceValue,
			iterator.GetValue(),
		)
	}

	resValue.Elem().Set(sliceValue)
}

func (m enumerable) valueToMap(resValue reflect.Value) {
	iterator := m.GetEnumerator()
	mapValue := resValue.Elem()
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		mapValue.SetMapIndex(
			iterator.GetKey(),
			iterator.GetValue(),
		)
	}
	resValue.Elem().Set(mapValue)
}
