package underscore

import (
	"reflect"
)

func Select(source, predicate interface{}) interface{} {
	var arrRV reflect.Value
	each(source, predicate, func (okRV, valueRV, _ reflect.Value) bool {
		if okRV.Bool() {
			if !arrRV.IsValid() {
				arrRT := reflect.SliceOf(valueRV.Type())
				arrRV = reflect.MakeSlice(arrRT, 0, 0)
			}

			arrRV = reflect.Append(arrRV, valueRV)
		}
		return false
	})
	if arrRV.IsValid() {
		return arrRV.Interface()
	}

	return nil
}

func SelectBy(source interface{}, properties map[string]interface{}) interface{} {
	return Select(source, func (value, _ interface{}) bool {
		return IsMatch(value, properties)
	})
}

//# chain
func (this *Query) Select(predicate interface{}) Queryer {
	this.source = Select(this.source, predicate)
	return this
}

func (this *Query) SelectBy(properties map[string]interface{}) Queryer {
	this.source = SelectBy(this.source, properties)
	return this
}