package underscore

import (
	"reflect"
)

func First(source interface{}) interface{} {
	var rv reflect.Value
	each(source, func (value, _ interface{}) bool {
		rv = reflect.ValueOf(value)
		return rv.IsValid()
	}, nil)
	if rv.IsValid() {
		return rv.Interface()
	}

	return nil
}

//# chain
func (this *Query) First() Queryer {
	this.source = First(this.source)
	return this
}