package underscore

import "reflect"

var (
	facadeRT      = reflect.TypeOf(facade{})
	nilEnumerable = enumerable{
		Enumerator: func() IEnumerator {
			return nullEnumerator{
				Src: nilRV,
			}
		},
	}
	nilRV     = reflect.ValueOf(nil)
	rtOfNilRV = reflect.TypeOf(nilRV)
)
