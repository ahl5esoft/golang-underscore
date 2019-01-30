package underscore

import (
	"reflect"
)

func (m *query) Value(v interface{}) {
	if m.Source == nil {
		return
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

	rv.Elem().Set(
		reflect.ValueOf(m.Source),
	)
}
