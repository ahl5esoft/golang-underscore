package underscore

import (
	"errors"
	"reflect"
)

var EMPTY_ARRAY = make([]interface{}, 0)

func Map(source interface{}, selector func(interface{}, interface{}) interface{}) ([]interface{}, error) {
	if selector == nil {
		return EMPTY_ARRAY, errors.New("underscore: Map's selector is nil")
	}

	if source == nil {
		return EMPTY_ARRAY, nil
	}

	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			if sourceRV.Len() == 0 {
				return EMPTY_ARRAY, nil
			}

			results := make([]interface{}, sourceRV.Len())
			for i := 0; i < sourceRV.Len(); i++ {
				results[i] = selector(
					sourceRV.Index(i).Interface(),
					i,
				)
			}
			return results, nil
		case reflect.Map:
			keyRVs := sourceRV.MapKeys()
			if len(keyRVs) == 0 {
				return EMPTY_ARRAY, nil
			}

			results := make([]interface{}, len(keyRVs))
			for i := 0; i < len(keyRVs); i++ {
				results[i] = selector(
					sourceRV.MapIndex(keyRVs[i]).Interface(),
					keyRVs[i].Interface(),
				)
			}
			return results, nil
	}
	return EMPTY_ARRAY, nil
}

//chain
func (this *Query) Map(selector func(interface{}, interface{}) interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = Map(this.source, selector)
	}
	return this
}