package underscore

import (
	"reflect"
	"sort"
)

// Sort is 排序
func Sort(source, selector, result interface{}) {
	resultRV := reflect.ValueOf(result)
	if resultRV.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

	qs := sortQuery{}
	qs.Sort(source, selector)
	if qs.Len() == 0 {
		return
	}

	sort.Sort(qs)
	resultRV.Elem().Set(qs.ValuesRV)
}

// SortBy is 根据属性排序
func SortBy(source interface{}, property string, result interface{}) {
	getPropertyRV := PropertyRV(property)
	Sort(source, func(value, _ interface{}) facade {
		return facade{
			getPropertyRV(value),
		}
	}, result)
}

func (m *query) Sort(selector interface{}) IQuery {
	Sort(m.Source, selector, &m.Source)
	return m
}

func (m *query) SortBy(property string) IQuery {
	SortBy(m.Source, property, &m.Source)
	return m
}
