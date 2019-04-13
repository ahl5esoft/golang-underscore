package underscore

import (
	"sort"
)

// Sort is 排序
func Sort(source, selector interface{}) interface{} {
	qs := sortQuery{}
	qs.Sort(source, selector)
	if qs.Len() == 0 {
		return nil
	}

	sort.Sort(qs)
	return qs.ValuesRV.Interface()
}

// SortBy is 根据属性排序
func SortBy(source interface{}, property string) interface{} {
	getPropertyRV := PropertyRV(property)
	return Sort(source, func(value, _ interface{}) facade {
		return facade{
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
