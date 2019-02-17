package underscore

import (
	"reflect"
	"sort"
)

// Reverse is 倒序
func Reverse(source, selector, result interface{}) {
	rv := reflect.ValueOf(result)
	if rv.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

	qs := sortQuery{}
	qs.Sort(source, selector)
	if qs.Len() == 0 {
		return
	}

	sort.Sort(sort.Reverse(qs))
	rv.Elem().Set(qs.ValuesRV)
}

// ReverseBy is 根据属性倒序
func ReverseBy(source interface{}, property string, result interface{}) {
	getPropertyRV := PropertyRV(property)
	Reverse(source, func(value, _ interface{}) facade {
		return facade{
			getPropertyRV(value),
		}
	}, result)
}

func (m *query) Reverse(selector interface{}) IQuery {
	Reverse(m.Source, selector, &m.Source)
	return m
}

func (m *query) ReverseBy(property string) IQuery {
	ReverseBy(m.Source, property, &m.Source)
	return m
}
