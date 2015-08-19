package underscore

import (
	"reflect"
)

var EMPTY_MAP = make(map[interface{}]interface{})

func Index(source interface{}, indexSelector func(interface{}) (interface{}, error)) (map[interface{}]interface{}, error) {
	if source == nil {
		return EMPTY_MAP, nil
	}

	dict := make(map[interface{}]interface{})
	fnRV := reflect.ValueOf(indexSelector)
	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			if sourceRV.Len() == 0 {
				return dict, nil
			}

			for i := 0; i < sourceRV.Len(); i++ {
				keyRVs := fnRV.Call([]reflect.Value{
					sourceRV.Index(i),
				})
				if !keyRVs[1].IsNil() {
					return EMPTY_MAP, keyRVs[1].Interface().(error)	
				}

				dict[ToInterface(keyRVs[0])] = ToInterface(sourceRV.Index(i))
			}
			break
		case reflect.Map:
			oldKeyRVs := sourceRV.MapKeys()
			if len(oldKeyRVs) == 0 {
				return dict, nil
			}

			for i := 0; i < len(oldKeyRVs); i++ {
				keyRVs := fnRV.Call([]reflect.Value{
					sourceRV.MapIndex(oldKeyRVs[i]),
				})
				if !keyRVs[1].IsNil() {
					return EMPTY_MAP, keyRVs[1].Interface().(error)	
				}

				dict[ToInterface(keyRVs[0])] = ToInterface(sourceRV.MapIndex(oldKeyRVs[i]))
			}
			break
	}
	return dict, nil
}

func IndexBy(source interface{}, field string) (map[interface{}]interface{}, error) {
	return Index(source, func (item interface{}) (interface{}, error) {
		return getFieldValue(item, field)
	})
}

//Chain
func (this *Query) Index(indexSelector func(item interface{}) (interface{}, error)) Queryer {
	if this.err != nil {
		return this
	}

	this.source, this.err = Index(this.source, indexSelector)
	return this
}

func (this *Query) IndexBy(field string) Queryer {
	if this.err != nil {
		return this
	}

	this.source, this.err = IndexBy(this.source, field)
	return this
}