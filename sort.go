package underscore

import (
	"reflect"
	"sort"
)

// Sort is 排序
func Sort(source, selector interface{}) interface{} {
	qs := sortQuery{}
	each(source, selector, func(sortRV, valueRV, _ reflect.Value) bool {
		if qs.Len() == 0 {
			keysRT := reflect.SliceOf(sortRV.Type())
			qs.KeysRV = reflect.MakeSlice(keysRT, 0, 0)

			valuesRT := reflect.SliceOf(valueRV.Type())
			qs.ValuesRV = reflect.MakeSlice(valuesRT, 0, 0)
		}

		qs.KeysRV = reflect.Append(qs.KeysRV, sortRV)
		qs.ValuesRV = reflect.Append(qs.ValuesRV, valueRV)
		return false
	})
	if qs.Len() > 0 {
		sort.Sort(qs)
		return qs.ValuesRV.Interface()
	}

	return nil
}

// SortBy is 根据属性排序
func SortBy(source interface{}, property string) interface{} {
	getPropertyRV := PropertyRV(property)
	return Sort(source, func(value, _ interface{}) Facade {
		return Facade{
			getPropertyRV(value),
		}
	})
}

func (m *query) Sort(selector interface{}) IQuery {
	m.Source = Sort(m.Source, selector)
	return m
}

func (m *query) SortBy(property string) IQuery {
	m.Source = SortBy(m.Source, property)
	return m
}
