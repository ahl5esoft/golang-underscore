package underscore

import (
	"reflect"
)

// ToRealValue is 将反射值转为真实类型的值
func ToRealValue(rv reflect.Value) interface{} {
	var value interface{}
	switch rv.Kind() {
	case reflect.Bool:
		value = rv.Bool()
		break
	case reflect.Float32, reflect.Float64:
		value = rv.Float()
		break
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		value = rv.Int()
		break
	case reflect.String:
		value = rv.String()
		break
	case reflect.Struct:
		value = rv.Interface()
		break
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value = rv.Uint()
		break
	case reflect.Ptr:
		return ToRealValue(
			reflect.Indirect(rv),
		)
	default:
		if !rv.IsNil() {
			value = rv.Interface()
		}
		break
	}
	return value
}
