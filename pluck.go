package underscore

import (
	"reflect"
)

func Pluck(source interface{}, property string) (interface{}, error) {
	var mapRV reflect.Value
	err := each(source, func (args []reflect.Value) (bool, reflect.Value) {
		pRV, err := getPropertyRV(args[0], property)

		if err == nil {
			if !mapRV.IsValid() {
				mapRV = makeSliceRVWithElem(pRV.Type(), 0)
			}

			mapRV = reflect.Append(mapRV, pRV)
		}
		return false, reflect.ValueOf(err)
	})
	if err == nil && mapRV.IsValid() {
		return mapRV.Interface(), nil
	}

	return nil, err
}

//chain
func (this *Query) Pluck(property string) Queryer {
	if this.err == nil {
		this.source, this.err = Pluck(this.source, property)
	}
	return this
}