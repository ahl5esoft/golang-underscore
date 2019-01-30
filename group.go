package underscore

import (
	"reflect"
)

// Group is 分组
func Group(source, keySelector, result interface{}) {
	rv := reflect.ValueOf(result)
	if rv.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

	var groupRV reflect.Value
	each(source, keySelector, func(groupKeyRV, valueRV, _ reflect.Value) bool {
		groupValueRT := reflect.SliceOf(valueRV.Type())
		if !groupRV.IsValid() {
			groupRT := reflect.MapOf(groupKeyRV.Type(), groupValueRT)
			groupRV = reflect.MakeMap(groupRT)
		}

		valuesRV := groupRV.MapIndex(groupKeyRV)
		if !valuesRV.IsValid() {
			valuesRV = reflect.MakeSlice(groupValueRT, 0, 0)
		}
		valuesRV = reflect.Append(valuesRV, valueRV)

		groupRV.SetMapIndex(groupKeyRV, valuesRV)
		return false
	})
	if groupRV.IsValid() {
		rv.Elem().Set(groupRV)
	}
}

// GroupBy is 根据某个属性分组
func GroupBy(source interface{}, property string, result interface{}) {
	getPropertyRV := PropertyRV(property)
	Group(source, func(value, _ interface{}) Facade {
		return Facade{
			getPropertyRV(value),
		}
	}, result)
}

func (m *query) Group(keySelector interface{}) IQuery {
	Group(m.Source, keySelector, &m.Source)
	return m
}

func (m *query) GroupBy(property string) IQuery {
	GroupBy(m.Source, property, &m.Source)
	return m
}
