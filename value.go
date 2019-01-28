package underscore

import (
	"fmt"
	"reflect"
)

func (m *query) Value(v interface{}) {
	if m.Source == nil {
		return
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		panic(
			fmt.Sprintf("receive type must be a pointer: `Chain.Value(v)`"),
		)
	}

	rv.Elem().Set(
		reflect.ValueOf(m.Source),
	)
}
