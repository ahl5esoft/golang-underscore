package underscore

import (
	"errors"
	"reflect"
)

var EMPTY_GROUP = make(map[interface{}][]interface{})

func Group(source, keySelector interface{}) (interface{}, error) {
	ksRV := reflect.ValueOf(keySelector)
	if ksRV.Kind() != reflect.Func {
		return nil, errors.New("underscore: Group's keySelector is not func")
	}

	if source == nil {
		return nil, nil
	}
	
	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			if sourceRV.Len() == 0 {
				return nil, nil
			}

			groupRV := makeMapRV(ksRV.Type().Out(0), sourceRV.Type())
			for i := 0; i < sourceRV.Len(); i++ {
				values := ksRV.Call(
					[]reflect.Value{
						sourceRV.Index(i),
						reflect.ValueOf(i),
					},
				)
				if !values[1].IsNil() {
					return nil, values[1].Interface().(error)
				}

				valuesRV := groupRV.MapIndex(values[0])
				if !valuesRV.IsValid() {
					valuesRV = makeSliceRV(sourceRV.Type())
				}
				valuesRV = reflect.Append(valuesRV, sourceRV.Index(i))
				
				groupRV.SetMapIndex(values[0], valuesRV)
			}
			return groupRV.Interface(), nil
		case reflect.Map:
			keyRVs := sourceRV.MapKeys()
			if len(keyRVs) == 0 {
				return nil, nil
			}

			groupRV := makeGroupRV(ksRV.Type().Out(0), sourceRV.MapIndex(keyRVs[0]).Type())
			for _, keyRV := range keyRVs {
				values := ksRV.Call(
					[]reflect.Value{
						sourceRV.MapIndex(keyRV),
						keyRV,
					},
				)
				if !values[1].IsNil() {
					return nil, values[1].Interface().(error)
				}

				valuesRV := groupRV.MapIndex(values[0])
				if !valuesRV.IsValid() {
					valuesRV = makeSliceRVWithElem(sourceRV.MapIndex(keyRV).Type())
				}
				valuesRV = reflect.Append(valuesRV, sourceRV.MapIndex(keyRV))
				
				groupRV.SetMapIndex(values[0], valuesRV)
			}
	}
	return nil, nil
}

func GroupBy(source interface{}, property string) (interface{}, error) {
	return Group(source, func (item, _ interface{}) (interface{}, error) {
		return getPropertyValue(item, property)
	})
}

//Chain
func (this *Query) Group(keySelector interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = Group(this.source, keySelector)
	}
	return this
}

func (this *Query) GroupBy(property string) Queryer {
	if this.err == nil {
		this.source, this.err = GroupBy(this.source, property)
	}
	return this
}