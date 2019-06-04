package underscore

import (
	"reflect"
)

func filter(source, predicate interface{}, compareValue bool) interface{} {
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
		return tempRV.Interface()
	}

	return nil
}

func (m enumerable) Filter(predicate interface{}) IEnumerable {
	return m.Where(predicate)
}

func (m enumerable) FilterBy(dict map[string]interface{}) IEnumerable {
	return m.WhereBy(dict)
}
