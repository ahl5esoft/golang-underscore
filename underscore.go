package underscore

import (
	"errors"
	"reflect"
)

// Facade is 包装真实的值 each的时候 如果iterator返回Facade 则会将Real替换iterator返回值
type Facade struct {
	Real reflect.Value
}

var (
	// ErrorRt is 错误类型
	ErrorRt = reflect.TypeOf(errors.New(""))
	// FacadeRt is 门面类型
	FacadeRt = reflect.TypeOf(Facade{})
	// NullRv is 反射值
	NullRv = reflect.ValueOf(nil)
	// NullRvOfRt is nil反射值类型
	NullRvOfRt = reflect.TypeOf(NullRv)
)
