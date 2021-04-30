package underscore

import (
	"reflect"
	"strings"
)

// GetFieldValueFunc is get field reflect.Value func
type GetFieldValueFunc func(interface{}) reflect.Value

// Field is 获取字段函数
func Field(name string) func(interface{}) interface{} {
	fn := FieldValue(name)
	return func(item interface{}) interface{} {
		return fn(item).Interface()
	}
}

// FieldValue is 获取reflect.Value
func FieldValue(name string) GetFieldValueFunc {
	var getter GetFieldValueFunc
	getter = func(item interface{}) reflect.Value {
		itemRV := getRealValue(item)
		if itemRV.Kind() == reflect.Ptr {
			itemRV = itemRV.Elem()
		}

		itemRT := itemRV.Type()
		for i := 0; i < itemRT.NumField(); i++ {
			field := itemRT.Field(i)
			if field.Anonymous {
				rv := getter(
					itemRV.Field(i),
				)
				if rv != nilValue {
					return rv
				}
			}

			if strings.EqualFold(name, field.Name) {
				return itemRV.Field(i)
			}
		}

		return nilValue
	}
	return getter
}
