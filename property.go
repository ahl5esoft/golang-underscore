package underscore

import (
	"reflect"
	"strings"
)

// GetProeprtyRVFunc is get property reflect.Value func
type GetProeprtyRVFunc func(interface{}) reflect.Value

// Property is 获取属性函数
func Property(name string) func(interface{}) interface{} {
	fn := PropertyRV(name)
	return func(item interface{}) interface{} {
		return fn(item).Interface()
	}
}

// PropertyRV is 获取reflect.Value
func PropertyRV(name string) GetProeprtyRVFunc {
	var getter GetProeprtyRVFunc
	getter = func(item interface{}) reflect.Value {
		itemRV := getRealRV(item)
		itemRT := itemRV.Type()
		for i := 0; i < itemRT.NumField(); i++ {
			field := itemRT.Field(i)
			if field.Anonymous {
				rv := getter(
					itemRV.Field(i),
				)
				if rv != nilRV {
					return rv
				}
			}

			if strings.ToLower(name) == strings.ToLower(field.Name) {
				return itemRV.Field(i)
			}
		}

		return nilRV
	}
	return getter
}
