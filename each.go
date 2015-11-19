package underscore

import (
	"errors"
	"reflect"
)

func Each(source, iterator interface{}) error {
	iteratorRV := reflect.ValueOf(iterator)
	if iteratorRV.Kind() != reflect.Func {
		return errors.New("underscore: Each's iterator is not func")
	}

	return each(source, func (args []reflect.Value) (bool, reflect.Value) {
		values := iteratorRV.Call(args)
		return false, values[0]
	})
}

//Chain
func (this *Query) Each(iterator interface{}) Queryer {
	if this.err == nil {
		this.err = Each(this.source, iterator)
	}
	return this
}