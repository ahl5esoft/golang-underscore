package underscore

import (
	"reflect"
)

// Index is 转化为indexSelector筛选出的值为key的map
func Index(source, indexSelector interface{}) interface{} {
	var tempRV reflect.Value
	each(source, indexSelector, func(indexRV, valueRV, _ reflect.Value) bool {
		if !tempRV.IsValid() {
			tempRT := reflect.MapOf(indexRV.Type(), valueRV.Type())
			tempRV = reflect.MakeMap(tempRT)
		}

		tempRV.SetMapIndex(indexRV, valueRV)
		return false
	})
	if tempRV.IsValid() {
		return tempRV.Interface()
	}
	return nil
}

// IndexBy is 转化为property值的map
func IndexBy(source interface{}, property string) interface{} {
	getPropertyRV := PropertyRV(property)
	return Index(source, func(value, _ interface{}) facade {
		return facade{
			getPropertyRV(value),
		}
	})
}

func (m *query) Index(indexSelector interface{}) IQuery {
	m.Source = Index(m.Source, indexSelector)
	return m
}

func (m *query) IndexBy(property string) IQuery {
	m.Source = IndexBy(m.Source, property)
	return m
}

func (m enumerable) Index(keySelector interface{}) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			iterator := m.GetEnumerator()
			keySelectorRV := reflect.ValueOf(keySelector)
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					if ok = iterator.MoveNext(); ok {
						keyRV = getFuncReturnRV(keySelectorRV, iterator)
						valueRV = iterator.GetValue()
					}

					return
				},
			}
		},
	}
}

func (m enumerable) IndexBy(fieldName string) IEnumerable {
	getter := PropertyRV(fieldName)
	return m.Index(func(value, _ interface{}) facade {
		return facade{
			getter(value),
		}
	})
}
