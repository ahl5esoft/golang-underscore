package underscore

import (
	"reflect"
)

// Group is 分组
func Group(source, keySelector interface{}) interface{} {
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
		return groupRV.Interface()
	}

	return nil
}

// GroupBy is 根据某个属性分组
func GroupBy(source interface{}, property string) interface{} {
	getPropertyRV := PropertyRV(property)
	return Group(source, func(value, _ interface{}) Facade {
		return Facade{
			getPropertyRV(value),
		}
	})
}

func (m *query) Group(keySelector interface{}) IQuery {
	m.Source = Group(m.Source, keySelector)
	return m
}

func (m *query) GroupBy(property string) IQuery {
	m.Source = GroupBy(m.Source, property)
	return m
}
