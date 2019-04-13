package underscore

import (
	"sort"
)

// Reverse is 倒序
func Reverse(source, selector interface{}) interface{} {
	qs := sortQuery{}
	qs.Sort(source, selector)
	if qs.Len() == 0 {
		return nil
	}

	sort.Sort(sort.Reverse(qs))
	return qs.ValuesRV.Interface()
}

// ReverseBy is 根据属性倒序
func ReverseBy(source interface{}, property string) interface{} {
	getPropertyRV := PropertyRV(property)
	return Reverse(source, func(value, _ interface{}) facade {
		return facade{
			getPropertyRV(value),
		}
	})
}

func (m *query) Reverse(selector interface{}) IQuery {
	m.Source = Reverse(m.Source, selector)
	return m
}

func (m *query) ReverseBy(property string) IQuery {
	m.Source = ReverseBy(m.Source, property)
	return m
}
