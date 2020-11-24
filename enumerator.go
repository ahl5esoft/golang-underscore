package underscore

import (
	"reflect"
)

type enumerator struct {
	MoveNextFunc func() (reflect.Value, reflect.Value, bool)

	key   reflect.Value
	value reflect.Value
}

func (m enumerator) GetKey() reflect.Value {
	return m.key
}

func (m enumerator) GetValue() reflect.Value {
	return getRealValue(m.value)
}

func (m *enumerator) MoveNext() (ok bool) {
	m.value, m.key, ok = m.MoveNextFunc()
	return
}
