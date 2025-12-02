package underscore

import "reflect"

type sorter struct {
	KeysValue   reflect.Value
	ValuesValue reflect.Value
	multiple    bool
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
	if m.multiple {
		// 多个字段排序
		for k := 0; k < m.KeysValue.Len(); k++ {
			thisRV := m.KeysValue.Index(k).Index(i)
			thatRV := m.KeysValue.Index(k).Index(j)
			
			result := compareValues(thisRV, thatRV)
			if result != 0 {
				return result < 0
			}
		}
		return false
	} else {
		// 单个字段排序（保持向后兼容）
		thisRV := m.KeysValue.Index(i)
		thatRV := m.KeysValue.Index(j)
		return compareValues(thisRV, thatRV) < 0
	}
}

func compareValues(thisRV, thatRV reflect.Value) int {
	switch thisRV.Kind() {
	case reflect.Float32, reflect.Float64:
		thisVal := thisRV.Float()
		thatVal := thatRV.Float()
		if thisVal < thatVal {
			return -1
		} else if thisVal > thatVal {
			return 1
		}
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		thisVal := thisRV.Int()
		thatVal := thatRV.Int()
		if thisVal < thatVal {
			return -1
		} else if thisVal > thatVal {
			return 1
		}
	case reflect.String:
		thisVal := thisRV.String()
		thatVal := thatRV.String()
		if thisVal < thatVal {
			return -1
		} else if thisVal > thatVal {
			return 1
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		thisVal := thisRV.Uint()
		thatVal := thatRV.Uint()
		if thisVal < thatVal {
			return -1
		} else if thisVal > thatVal {
			return 1
		}
	}
	return 0
}

func (m *sorter) Sort(iterator IEnumerator, selector interface{}) {
	m.multiple = false
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

func (m *sorter) SortMultiple(iterator IEnumerator, selectors []interface{}) {
	m.multiple = true
	selectorValues := make([]reflect.Value, len(selectors))
	for i, selector := range selectors {
		selectorValues[i] = reflect.ValueOf(selector)
	}
	
	// 收集所有元素
	var elements []reflect.Value
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		elements = append(elements, iterator.GetValue())
	}
	
	// 创建键切片的切片
	m.KeysValue = reflect.MakeSlice(reflect.SliceOf(reflect.SliceOf(reflect.TypeOf(interface{}{}))), len(selectors), len(selectors))
	
	// 为每个元素生成键
	for i, element := range elements {
		// 创建临时迭代器用于获取键值
		iter := &enumerator{
			MoveNextFunc: func() (valueValue reflect.Value, keyValue reflect.Value, ok bool) {
				ok = true
				valueValue = element
				keyValue = reflect.ValueOf(i)
				return
			},
		}
		
		// 为每个选择器生成键
		for j, selectorValue := range selectorValues {
			keyValue := getReturnValue(selectorValue, iter)
			
			if m.KeysValue.Index(j).IsZero() {
				// 初始化键切片
				keysType := reflect.SliceOf(keyValue.Type())
				m.KeysValue.Index(j).Set(reflect.MakeSlice(keysType, 0, 0))
			}
			
			m.KeysValue.Index(j).Set(reflect.Append(m.KeysValue.Index(j), keyValue))
		}
	}
	
	// 设置值切片
	if len(elements) > 0 {
		valuesType := reflect.SliceOf(elements[0].Type())
		m.ValuesValue = reflect.MakeSlice(valuesType, 0, 0)
		for _, element := range elements {
			m.ValuesValue = reflect.Append(m.ValuesValue, element)
		}
	}
}
