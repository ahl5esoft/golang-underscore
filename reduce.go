package underscore

import (
	"errors"
	"reflect"
)

func Reduce(source interface{}, iterator func(memo, value, key interface{}) interface{}, memo interface{}) (interface{}, error) {
	if iterator == nil {
		return memo, errors.New("underscore: Reduce's iterator is nil")
	}

	if source == nil {
		return memo, nil
	}

	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			if sourceRV.Len() == 0 {
				return memo, nil
			}

			for i := 0; i < sourceRV.Len(); i++ {
				memo = iterator(
					memo, 
					sourceRV.Index(i).Interface(), 
					i,
				)
			}
		case reflect.Map:
			oldKeyRVs := sourceRV.MapKeys()
			if len(oldKeyRVs) == 0 {
				return memo, nil
			}

			for i := 0; i < len(oldKeyRVs); i++ {
				memo = iterator(
					memo, 
					sourceRV.MapIndex(oldKeyRVs[i]).Interface(), 
					oldKeyRVs[i].Interface(),
				)
			}
	}
	return memo, nil
}

//Chain
func (this *Query) Reduce(iterator func(memo, value, key interface{}) interface{}, memo interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = Reduce(this.source, iterator, memo)
	}
	return this
}