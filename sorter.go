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
		
		var less bool
		equal := false
		
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
			return false
		}
		
		if !equal {
			return less
		}
	}
	
	// All fields are equal
	return false
}

func (m *sorter) Sort(iterator IEnumerator, selectors ...interface{}) {
	selectorValues := make([]reflect.Value, len(selectors))
	for i, selector := range selectors {
		selectorValues[i] = reflect.ValueOf(selector)
	}
	
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		if m.Len() == 0 {
			m.KeysValues = make([]reflect.Value, len(selectors))
			for i, selectorValue := range selectorValues {
				keyValue := getReturnValue(selectorValue, iterator)
				keysType := reflect.SliceOf(keyValue.Type())
				m.KeysValues[i] = reflect.MakeSlice(keysType, 0, 0)
			}
			
			valuesType := reflect.SliceOf(iterator.GetValue().Type())
			m.ValuesValue = reflect.MakeSlice(valuesType, 0, 0)
		}
		
		for i, selectorValue := range selectorValues {
			keyValue := getReturnValue(selectorValue, iterator)
			m.KeysValues[i] = reflect.Append(m.KeysValues[i], keyValue)
		}
		
		m.ValuesValue = reflect.Append(
			m.ValuesValue,
			iterator.GetValue(),
		)
	}
}
