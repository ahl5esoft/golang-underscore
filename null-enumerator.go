package underscore

import (
	"reflect"
)

type nullEnumerator struct {
	Src reflect.Value
}

func (m nullEnumerator) GetKey() reflect.Value {
	return nilValue
}

func (m nullEnumerator) GetValue() reflect.Value {
	if m.Src.IsValid() && m.Src.Type() == facadeType {
		return m.Src.Interface().(facade).Real
	}

	return m.Src
}

func (m nullEnumerator) MoveNext() bool {
	return false
}
