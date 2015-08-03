package underscore

import (
	"reflect"
)

func Index(source interface{}, indexSelector func(interface{}) interface{}) map[interface{}]interface{} {
	dict := make(map[interface{}]interface{})
	if source == nil {
		return dict
	}

	fnRV := reflect.ValueOf(indexSelector)
	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			for i := 0; i < sourceRV.Len(); i++ {
				keyRVs := fnRV.Call([]reflect.Value{
					sourceRV.Index(i),
				})
				dict[ToInterface(keyRVs[0])] = ToInterface(sourceRV.Index(i))
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
				dict[ToInterface(keyRVs[0])] = ToInterface(sourceRV.MapIndex(oldKeyRVs[i]))
			}
			break
	}
	return dict
}

func IndexBy(source interface{}, field string) map[interface{}]interface{} {
	return Index(source, func (item interface{}) interface{} {
		return ToInterface(reflect.ValueOf(item).FieldByName(field))
	})
}

//Chain
func (this *Query) Index(indexSelector func(item interface{}) interface{}) Queryer {
	this.source = Index(this.source, indexSelector)
	return this
}

func (this *Query) IndexBy(field string) Queryer {
	this.source = IndexBy(this.source, field)
	return this
}