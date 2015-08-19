package underscore

import (
	"reflect"
)

var EMPTY_GROUP = make(map[interface{}][]interface{})

func Group(source interface{}, keySelector func(interface{}) (interface{}, error)) (map[interface{}][]interface{}, error) {
	if source == nil {
		return EMPTY_GROUP, nil
	}

	dict := make(map[interface{}][]interface{})
	fnRV := reflect.ValueOf(keySelector)
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
					return EMPTY_GROUP, keyRVs[1].Interface().(error)
				}

				key := ToInterface(keyRVs[0])
				dict[key] = append(dict[key], ToInterface(sourceRV.Index(i)))
			}
			break
		case reflect.Map:
			oldKeyRVs := sourceRV.MapKeys()
			if len(oldKeyRVs) == 0 {
				return dict, nil
			}

			for i := 0; i < len(oldKeyRVs); i++ {
				keyRVs := fnRV.Call([]reflect.Value{
					sourceRV.Index(i),
				})
				if !keyRVs[1].IsNil() {
					return EMPTY_GROUP, keyRVs[1].Interface().(error)
				}

				key := ToInterface(keyRVs[0])
				dict[key] = append(dict[key], ToInterface(sourceRV.MapIndex(oldKeyRVs[i])))
			}
			break
	}
	return dict, nil
}

func GroupBy(source interface{}, field string) (map[interface{}][]interface{}, error) {
	return Group(source, func (item interface{}) (interface{}, error) {
		return getFieldValue(item, field)
	})
}

//Chain
func (this *Query) Group(keySelector func(item interface{}) (interface{}, error)) Queryer {
	if this.err != nil {
		return this
	}

	this.source, this.err = Group(this.source, keySelector)
	return this
}

func (this *Query) GroupBy(field string) Queryer {
	if this.err != nil {
		return this
	}

	this.source, this.err = GroupBy(this.source, field)
	return this
}