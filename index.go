package underscore

import (
	"errors"
	"reflect"
)

var EMPTY_MAP = make(map[interface{}]interface{})

func Index(source interface{}, indexSelector func(interface{}, interface{}) (interface{}, error)) (map[interface{}]interface{}, error) {
	if indexSelector == nil {
		return EMPTY_MAP, errors.New("underscore: Index's indexSelector is nil")
	}

	if source == nil {
		return EMPTY_MAP, nil
	}

	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			if sourceRV.Len() == 0 {
				return EMPTY_MAP, nil
			}

			dict := make(map[interface{}]interface{})
			for i := 0; i < sourceRV.Len(); i++ {
				value := sourceRV.Index(i).Interface()
				index, err := indexSelector(value, i)
				if err != nil {
					return EMPTY_MAP, err
				}

				dict[index] = value
			}
			return dict, nil
		case reflect.Map:
			oldKeyRVs := sourceRV.MapKeys()
			if len(oldKeyRVs) == 0 {
				return EMPTY_MAP, nil
			}

			dict := make(map[interface{}]interface{})
			for i := 0; i < len(oldKeyRVs); i++ {
				value := sourceRV.MapIndex(oldKeyRVs[i]).Interface()
				index, err := indexSelector(value, oldKeyRVs[i].Interface())
				if err != nil {
					return EMPTY_MAP, err
				}

				dict[index] = value
			}
			return dict, nil
	}
	return EMPTY_MAP, nil
}

func IndexBy(source interface{}, field string) (map[interface{}]interface{}, error) {
	return Index(source, func (item interface{}, _ interface{}) (interface{}, error) {
		return getFieldValue(item, field)
	})
}

//Chain
func (this *Query) Index(indexSelector func(interface{}, interface{}) (interface{}, error)) Queryer {
	if this.err == nil {
		this.source, this.err = Index(this.source, indexSelector)
	}
	return this
}

func (this *Query) IndexBy(field string) Queryer {
	if this.err == nil {
		this.source, this.err = IndexBy(this.source, field)
	}
	return this
}