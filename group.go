package underscore

import (
	"reflect"
)

func Group(source interface{}, keySelector func(interface{}) interface{}) map[interface{}][]interface{} {
	dict := make(map[interface{}][]interface{})
	if source == nil {
		return dict
	}

	fnRV := reflect.ValueOf(keySelector)
	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			for i := 0; i < sourceRV.Len(); i++ {
				keyRVs := fnRV.Call([]reflect.Value{
					sourceRV.Index(i),
				})
				key := ToInterface(keyRVs[0])
				dict[key] = append(dict[key], ToInterface(sourceRV.Index(i)))
			}
			break
		case reflect.Map:
			oldKeyRVs := sourceRV.MapKeys()
			if len(oldKeyRVs) == 0 {
				return dict
			}

			for i := 0; i < len(oldKeyRVs); i++ {
				keyRVs := fnRV.Call([]reflect.Value{
					sourceRV.MapIndex(oldKeyRVs[i]),
				})
				key := ToInterface(keyRVs[0])
				dict[key] = append(dict[key], ToInterface(sourceRV.MapIndex(oldKeyRVs[i])))
			}
			break
	}
	return dict
}

func GroupBy(source interface{}, field string) map[interface{}][]interface{} {
	return Group(source, func (item interface{}) interface{} {
		return ToInterface(reflect.ValueOf(item).FieldByName(field))
	})
}

//Chain
func (this *Query) Group(keySelector func(item interface{}) interface{}) Queryer {
	this.source = Group(this.source, keySelector)
	return this
}

func (this *Query) GroupBy(field string) Queryer {
	this.source = GroupBy(this.source, field)
	return this
}