package underscore

import (
	"errors"
	"reflect"
)

func Select(source, predicate interface{}) (interface{}, error) {
	predicateRV := reflect.ValueOf(predicate)
	if predicateRV.Kind() != reflect.Func {
		return nil, errors.New("underscore: Select's predicate is not func")
	}

	var arrRV reflect.Value
	err := each(source, func (args []reflect.Value) (bool, reflect.Value) {
		if !arrRV.IsValid() {
			arrRV = makeSliceRVWithElem(args[0].Type(), 0)
		}

		values := predicateRV.Call(args)
		if !isErrorRVValid(values[1]) && values[0].Bool() {			
			arrRV = reflect.Append(arrRV, args[0])
		}

		return false, values[1]
	})
	if err == nil && arrRV.IsValid() {
		return arrRV.Interface(), nil
	}

	return nil, err
}

func SelectBy(source interface{}, properties map[string]interface{}) (interface{}, error) {
	if source == nil || properties == nil || len(properties) == 0 {
		return nil, nil
	}

	return Select(source, func (item, _ interface{}) (bool, error) {
		return All(properties, func (pv interface{}, pn string) (bool, error) {
			value, err := getPropertyValue(item, pn)
			if err != nil {
				return false, err
			}

			return value == pv, nil
		})
	})
}

//# chain
func (this *Query) Select(predicate interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = Select(this.source, predicate)
	}
	return this
}

func (this *Query) SelectBy(properties map[string]interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = SelectBy(this.source, properties)
	}
	return this
}