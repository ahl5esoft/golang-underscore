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
	if m.value.Type() == facadeRT {
		return m.value.Interface().(facade).Real
	}

	return m.value
}

func (m *enumerator) MoveNext() (ok bool) {
	m.value, m.key, ok = m.MoveNextFunc()
	return
}
