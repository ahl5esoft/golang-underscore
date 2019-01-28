package underscore

import (
	"reflect"
)

// FindIndex is 根据断言函数获取下标
func FindIndex(source, predicate interface{}) int {
	index := -1

	if !IsArray(source) {
		return index
	}

	each(source, predicate, func(okRV, _, keyRV reflect.Value) bool {
		ok := okRV.Bool()
		if ok {
			index = int(keyRV.Int())
		}
		return ok
	})

	return index
}

// FindIndexBy is 根据字典获取下标
func FindIndexBy(source interface{}, properties map[string]interface{}) int {
	return FindIndex(source, func(item interface{}, _ int) bool {
		return IsMatch(item, properties)
	})
}

func (m *query) FindIndex(predicate interface{}) int {
	return FindIndex(m.Source, predicate)
}

func (m *query) FindIndexBy(properties map[string]interface{}) int {
	return FindIndexBy(m.Source, properties)
}
