package underscore

import "reflect"

// IEnumerator is 迭代器接口
type IEnumerator interface {
	GetKey() reflect.Value
	GetValue() reflect.Value
	MoveNext() bool
}
