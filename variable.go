package underscore

import "reflect"

var (
	facadeType    = reflect.TypeOf(facade{})
	nilEnumerable = enumerable{
		Enumerator: func() IEnumerator {
			return nullEnumerator{
				Src: nilValue,
			}
		},
	}
	nilValue  = reflect.ValueOf(nil)
	valueType = reflect.TypeOf(nilValue)
)
