package underscore

import "reflect"

type sorter struct {
	KeysValues  []reflect.Value
	ValuesValue reflect.Value
}

func (m sorter) Len() int {
	if len(m.KeysValues) > 0 && m.KeysValues[0].IsValid() {
		return m.KeysValues[0].Len()
	}

	return 0
}

func (m sorter) Swap(i, j int) {
	for _, keysValue := range m.KeysValues {
		temp := keysValue.Index(i).Interface()
		keysValue.Index(i).Set(
			keysValue.Index(j),
		)
		keysValue.Index(j).Set(
			reflect.ValueOf(temp),
		)
	}

	temp := m.ValuesValue.Index(i).Interface()
	m.ValuesValue.Index(i).Set(
		m.ValuesValue.Index(j),
	)
	m.ValuesValue.Index(j).Set(
		reflect.ValueOf(temp),
	)
}

func (m sorter) Less(i, j int) bool {
	for _, keysValue := range m.KeysValues {
		thisRV := keysValue.Index(i)
		thatRV := keysValue.Index(j)
		less := false
		equal := true

		switch thisRV.Kind() {
		case reflect.Float32, reflect.Float64:
			thisVal := thisRV.Float()
			thatVal := thatRV.Float()
			less = thisVal < thatVal
			equal = thisVal == thatVal
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			thisVal := thisRV.Int()
			thatVal := thatRV.Int()
			less = thisVal < thatVal
			equal = thisVal == thatVal
		case reflect.String:
			thisVal := thisRV.String()
			thatVal := thatRV.String()
			less = thisVal < thatVal
			equal = thisVal == thatVal
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			thisVal := thisRV.Uint()
			thatVal := thatRV.Uint()
			less = thisVal < thatVal
			equal = thisVal == thatVal
		default:
			less = false
			equal = false
		}

		if !equal {
			return less
		}
	}
	return false
}

func (m *sorter) Sort(iterator IEnumerator, selector interface{}) {
	selectorValue := reflect.ValueOf(selector)
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		keyValue := getReturnValue(selectorValue, iterator)
		value := iterator.GetValue()
		if m.Len() == 0 {
			keysType := reflect.SliceOf(keyValue.Type())
			m.KeysValues = []reflect.Value{reflect.MakeSlice(keysType, 0, 0)}
			m.KeysValues[0] = reflect.Append(m.KeysValues[0], keyValue)

			valuesType := reflect.SliceOf(value.Type())
			m.ValuesValue = reflect.MakeSlice(valuesType, 0, 0)
			m.ValuesValue = reflect.Append(m.ValuesValue, value)
		} else {
			m.KeysValues[0] = reflect.Append(m.KeysValues[0], keyValue)
			m.ValuesValue = reflect.Append(m.ValuesValue, value)
		}
	}
}

func (m *sorter) SortMany(iterator IEnumerator, selectors ...interface{}) {
	if len(selectors) == 0 {
		return
	}
	if len(selectors) == 1 {
		m.Sort(iterator, selectors[0])
		return
	}

	selectorValues := make([]reflect.Value, len(selectors))
	for i, selector := range selectors {
		selectorValues[i] = reflect.ValueOf(selector)
	}

	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		// 为每个selector获取keyValue
		keyValues := make([]reflect.Value, len(selectors))
		for i, selectorValue := range selectorValues {
			keyValues[i] = getReturnValue(selectorValue, iterator)
		}

		// 获取当前value
		value := iterator.GetValue()

		if m.Len() == 0 {
			// 初始化所有keys slice
			m.KeysValues = make([]reflect.Value, len(selectors))
			for i, keyValue := range keyValues {
				keysType := reflect.SliceOf(keyValue.Type())
				m.KeysValues[i] = reflect.MakeSlice(keysType, 0, 0)
				m.KeysValues[i] = reflect.Append(m.KeysValues[i], keyValue)
			}

			// 初始化values slice
			valuesType := reflect.SliceOf(value.Type())
			m.ValuesValue = reflect.MakeSlice(valuesType, 0, 0)
			m.ValuesValue = reflect.Append(m.ValuesValue, value)
		} else {
			// 为每个selector添加key
			for i, keyValue := range keyValues {
				m.KeysValues[i] = reflect.Append(m.KeysValues[i], keyValue)
			}

			// 添加value
			m.ValuesValue = reflect.Append(m.ValuesValue, value)
		}
	}
}
