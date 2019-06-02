package underscore

import "reflect"

// IsArray is 判断是否数组或者切片
func IsArray(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Array || rv.Kind() == reflect.Slice
}
