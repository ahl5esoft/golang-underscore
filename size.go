package underscore

import (
	"reflect"
)

func Size(source interface{}) int {
	length := 0
	if source == nil {
		return length
	}

	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			length = sourceRV.Len()
			break
		case reflect.Map:
			length = len(sourceRV.MapKeys())
			break
	}
	return length
}

//chain
func (this *Query) Size() Queryer {
	if this.err != nil {
		this.source = 0
	} else {
		this.source = Size(this.source)
	}
	return this
}