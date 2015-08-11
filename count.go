package underscore

import (
	"reflect"
)

func Count(source interface{}) int {
	count := 0
	if source == nil {
		return count
	}

	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			count = sourceRV.Len()
			break
		case reflect.Map:
			count = len(sourceRV.MapKeys())
			break
	}
	return count
}

//chain
func (this *Query) Count() Queryer {	
	this.source = Count(this.source)
	return this
}