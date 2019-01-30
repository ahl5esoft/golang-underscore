package underscore

import (
	"reflect"
)

type sortQuery struct {
	KeysRV   reflect.Value
	ValuesRV reflect.Value
}

func (m sortQuery) Len() int {
	if m.KeysRV.IsValid() {
		return m.KeysRV.Len()
	}

	return 0
}

func (m sortQuery) Swap(i, j int) {
	temp := m.KeysRV.Index(i).Interface()
	m.KeysRV.Index(i).Set(
		m.KeysRV.Index(j),
	)
	m.KeysRV.Index(j).Set(
		reflect.ValueOf(temp),
	)

	temp = m.ValuesRV.Index(i).Interface()
	m.ValuesRV.Index(i).Set(
		m.ValuesRV.Index(j),
	)
	m.ValuesRV.Index(j).Set(
		reflect.ValueOf(temp),
	)
}

func (m sortQuery) Less(i, j int) bool {
	thisRV := m.KeysRV.Index(i)
	thatRV := m.KeysRV.Index(j)
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

func (m *sortQuery) Sort(source, selector interface{}) {
	each(source, selector, func(sortRV, valueRV, _ reflect.Value) bool {
		if m.Len() == 0 {
			keysRT := reflect.SliceOf(sortRV.Type())
			m.KeysRV = reflect.MakeSlice(keysRT, 0, 0)

			valuesRT := reflect.SliceOf(valueRV.Type())
			m.ValuesRV = reflect.MakeSlice(valuesRT, 0, 0)
		}

		m.KeysRV = reflect.Append(m.KeysRV, sortRV)
		m.ValuesRV = reflect.Append(m.ValuesRV, valueRV)
		return false
	})
}
