package underscore

import (
	"reflect"
)

func filter(source, predicate interface{}, compareValue bool, result interface{}) {
	resultRV := reflect.ValueOf(result)
	if resultRV.Kind() != reflect.Ptr {
		panic("receive type must be a pointer")
	}

	var tempRV reflect.Value
	each(source, predicate, func(resRV, valueRV, _ reflect.Value) bool {
		if resRV.Bool() == compareValue {
			if !tempRV.IsValid() {
				tempRT := reflect.SliceOf(valueRV.Type())
				tempRV = reflect.MakeSlice(tempRT, 0, 0)
			}

			tempRV = reflect.Append(tempRV, valueRV)
		}
		return false
	})
	if tempRV.IsValid() {
		resultRV.Elem().Set(tempRV)
	}
}
