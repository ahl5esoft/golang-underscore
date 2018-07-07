package underscore

import (
	"reflect"

	fjson "github.com/json-iterator/go"
)

// Clone will create a deep-copied clone of the `source`
func Clone(source interface{}) interface{} {
	var (
		bytes  []byte
		cloned interface{}
	)

	rt := reflect.TypeOf(source)
	if rt.Kind() == reflect.Ptr {
		cloned = reflect.New(rt.Elem()).Interface()
		bytes, _ = fjson.Marshal(
			reflect.ValueOf(source).Elem().Interface(),
		)
	} else {
		cloned = reflect.New(rt).Interface()
		bytes, _ = fjson.Marshal(source)
	}

	fjson.Unmarshal(bytes, cloned)
	return reflect.ValueOf(cloned).Elem().Interface()
}

// Clone is Queryer's method
func (q *Query) Clone() Queryer {
	q.source = Clone(q.source)
	return q
}
