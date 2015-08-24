package underscore

import (
	"errors"
	"reflect"
)

func Uniq(source interface{}) ([]interface{}, error) {
	return uniq(source, nil), nil
}

func UniqBy(source interface{}, selector func(interface{}, int) interface{}) ([]interface{}, error) {
	if selector == nil {
		return EMPTY_ARRAY, errors.New("underscore: UniqBy's selector is nil")
	}

	return uniq(source, selector), nil
}

func uniq(source interface{}, selector func(interface{}, int) interface{}) []interface{} {
	if source == nil {
		return EMPTY_ARRAY
	}
	
	sourceRV := reflect.ValueOf(source)
	if sourceRV.Kind() == reflect.Array || sourceRV.Kind() == reflect.Slice {
		if sourceRV.Len() == 0 {
			return EMPTY_ARRAY
		}

		dict := make(map[interface{}]interface{})
		for i := 0; i < sourceRV.Len(); i++ {
			v := sourceRV.Index(i).Interface()
			var k interface{}
			if selector == nil {
				k = v
			} else {
				k = selector(v, i)
			}
			if _, ok := dict[k]; !ok {
				dict[k] = v
			}
		}

		results := make([]interface{}, len(dict))
		i := 0
		for k := range dict {
			results[i] = k
			i++
		}
		return results
	}
	return EMPTY_ARRAY
}

//chain
func (this *Query) Uniq() Queryer {
	if this.err == nil {
		this.source, this.err = Uniq(this.source)
	}
	return this
}

func (this *Query) UniqBy(selector func(interface{}, int) interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = UniqBy(this.source, selector)
	}
	return this
}