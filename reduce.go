package underscore

import (
	"errors"
	"reflect"
)

func Reduce(source interface{}, iterator func(memo, value, key interface{}) (interface{}, error), memo interface{}) (interface{}, error) {
	clone, _ := Clone(memo)
	if iterator == nil {
		return clone, errors.New("underscore: Reduce's iterator is nil")
	}

	if source == nil {
		return clone, nil
	}

	var err error
	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			if sourceRV.Len() == 0 {
				return clone, nil
			}

			for i := 0; i < sourceRV.Len(); i++ {
				memo, err = iterator(
					memo, 
					sourceRV.Index(i).Interface(), 
					i,
				)
				if err != nil {
					return clone, err
				}
			}
		case reflect.Map:
			oldKeyRVs := sourceRV.MapKeys()
			if len(oldKeyRVs) == 0 {
				return memo, nil
			}

			for i := 0; i < len(oldKeyRVs); i++ {
				memo, err = iterator(
					memo, 
					sourceRV.MapIndex(oldKeyRVs[i]).Interface(), 
					oldKeyRVs[i].Interface(),
				)
				if err != nil {
					return clone, err
				}
			}
	}
	return memo, nil
}

//Chain
func (this *Query) Reduce(iterator func(memo, value, key interface{}) (interface{}, error), memo interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = Reduce(this.source, iterator, memo)
	}
	return this
}