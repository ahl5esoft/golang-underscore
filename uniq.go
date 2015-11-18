package underscore

import (
	"errors"
	"reflect"
)

func Uniq(source, selector interface{}) (interface{}, error) {
	if selector == nil {
		selector = func (item, _ interface{}) (interface{}, error) {
			return item, nil
		}
	}

	selectorRV := reflect.ValueOf(selector)
	if selectorRV.Kind() != reflect.Func {
		return nil, errors.New("underscore: UniqBy's selector is not func")
	}

	var mapRV reflect.Value
	err := each(source, func (args []reflect.Value) (bool, reflect.Value) {
		if !mapRV.IsValid() {
			mapRV = makeMapRV(
				selectorRV.Type().Out(0),
				args[0].Type(),
			)
		}

		values := selectorRV.Call(args)
		if !isErrorRVValid(values[1]) {
			mapRV.SetMapIndex(
				values[0],
				args[0],
			)
		}

		return false, values[1]
	})
	if err == nil && mapRV.IsValid() {
		keyRVs := mapRV.MapKeys()
		arrRV := makeSliceRVWithElem(
			mapRV.MapIndex(keyRVs[0]).Type(),
			len(keyRVs),
		)
		for i := 0; i < len(keyRVs); i++ {
			arrRV.Index(i).Set(
				mapRV.MapIndex(keyRVs[i]),
			)
		}
		return arrRV.Interface(), nil
	}

	return nil, err
}

func UniqBy(source interface{}, property string) (interface{}, error) {
	return Uniq(source, func (item, _ interface{}) (interface{}, error) {
		return getPropertyValue(item, property)
	})
}

//chain
func (this *Query) Uniq(selector interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = Uniq(this.source, selector)
	}
	return this
}

func (this *Query) UniqBy(property string) Queryer {
	if this.err == nil {
		this.source, this.err = UniqBy(this.source, property)
	}
	return this
}