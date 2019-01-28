package underscore

import (
	"reflect"

	fjson "github.com/json-iterator/go"
)

// Clone will create a deep-copied clone of the `src`
func Clone(src, dst interface{}) {
	bf, _ := fjson.Marshal(src)
	fjson.Unmarshal(bf, dst)
}

func (m *query) Clone() IQuery {
	rt := reflect.TypeOf(m.Source)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	rv := reflect.New(rt)
	Clone(
		m.Source,
		rv.Interface(),
	)

	if rt.Kind() == reflect.Array || rt.Kind() == reflect.Map || rt.Kind() == reflect.Slice {
		m.Source = rv.Elem().Interface()
	} else {
		m.Source = rv.Interface()
	}
	return m
}
