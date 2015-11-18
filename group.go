package underscore

import (
	"errors"
	"reflect"
)

func Group(source, keySelector interface{}) (interface{}, error) {
	ksRV := reflect.ValueOf(keySelector)
	if ksRV.Kind() != reflect.Func {
		return nil, errors.New("underscore: Group's keySelector is not func")
	}

	var groupRV reflect.Value
	err := each(source, func (args []reflect.Value) (bool, reflect.Value) {
		if !groupRV.IsValid() {
			groupRV = makeGroupRV(ksRV.Type().Out(0), args[0].Type())
		}

		values := ksRV.Call(args)
		if !isErrorRVValid(values[1]) {
			valuesRV := groupRV.MapIndex(values[0])
			if !valuesRV.IsValid() {
				valuesRV = makeSliceRVWithElem(args[0].Type(), 0)
			}
			valuesRV = reflect.Append(valuesRV, args[0])
			
			groupRV.SetMapIndex(values[0], valuesRV)
		}

		return false, values[1]
	})
	if err == nil && groupRV.IsValid() {
		return groupRV.Interface(), nil
	}

	return nil, err
}

func GroupBy(source interface{}, property string) (interface{}, error) {
	var groupRV reflect.Value
	err := each(source, func (args []reflect.Value) (bool, reflect.Value) {
		pRV, err := getPropertyRV(args[0], property)
		if err == nil {
			if !groupRV.IsValid() {
				groupRV = makeGroupRV(pRV.Type(), args[0].Type())
			}

			valuesRV := groupRV.MapIndex(pRV)
			if !valuesRV.IsValid() {
				valuesRV = makeSliceRVWithElem(args[0].Type(), 0)
			}
			valuesRV = reflect.Append(valuesRV, args[0])
			
			groupRV.SetMapIndex(pRV, valuesRV)
		}
		return false, reflect.ValueOf(err)
	})
	if err == nil && groupRV.IsValid() {
		return groupRV.Interface(), nil
	}

	return nil, err
}

//Chain
func (this *Query) Group(keySelector interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = Group(this.source, keySelector)
	}
	return this
}

func (this *Query) GroupBy(property string) Queryer {
	if this.err == nil {
		this.source, this.err = GroupBy(this.source, property)
	}
	return this
}