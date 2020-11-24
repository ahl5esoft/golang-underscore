package underscore

import "reflect"

type sorter struct {
	KeysValue   reflect.Value
	ValuesValue reflect.Value
}

func (m sorter) Len() int {
	if m.KeysValue.IsValid() {
		return m.KeysValue.Len()
	}

	return 0
}

func (m sorter) Swap(i, j int) {
	temp := m.KeysValue.Index(i).Interface()
	m.KeysValue.Index(i).Set(
		m.KeysValue.Index(j),
	)
	m.KeysValue.Index(j).Set(
		reflect.ValueOf(temp),
	)

	temp = m.ValuesValue.Index(i).Interface()
	m.ValuesValue.Index(i).Set(
		m.ValuesValue.Index(j),
	)
	m.ValuesValue.Index(j).Set(
		reflect.ValueOf(temp),
	)
}

func (m sorter) Less(i, j int) bool {
	thisRV := m.KeysValue.Index(i)
	thatRV := m.KeysValue.Index(j)
	switch thisRV.Kind() {
	case reflect.Float32, reflect.Float64:
		return thisRV.Float() < thatRV.Float()
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return thisRV.Int() < thatRV.Int()
	case reflect.String:
		return thisRV.String() < thatRV.String()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return thisRV.Uint() < thatRV.Uint()
	default:
		return false
	}
}

func (m *sorter) Sort(iterator IEnumerator, selector interface{}) {
	selectorValue := reflect.ValueOf(selector)
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		keyValue := getReturnValue(selectorValue, iterator)
		if m.Len() == 0 {
			keysType := reflect.SliceOf(
				keyValue.Type(),
			)
			m.KeysValue = reflect.MakeSlice(keysType, 0, 0)

			valuesType := reflect.SliceOf(
				iterator.GetValue().Type(),
			)
			m.ValuesValue = reflect.MakeSlice(valuesType, 0, 0)
		}

		m.KeysValue = reflect.Append(m.KeysValue, keyValue)
		m.ValuesValue = reflect.Append(
			m.ValuesValue,
			iterator.GetValue(),
		)
	}
}
