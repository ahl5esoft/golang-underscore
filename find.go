package underscore

import (
	"errors"
	"reflect"
)

func Find(source, predicate interface{}) (interface{}, error) {
	predicateRV := reflect.ValueOf(predicate)
	if predicateRV.Kind() != reflect.Func {
		return nil, errors.New("underscore: Find's predicate is not func")
	}

	var res interface{}
	err := each(source, func (args []reflect.Value) (bool, reflect.Value) {
		values := predicateRV.Call(args)
		if !isErrorRVValid(values[1]) && values[0].Bool() {
			res = args[0].Interface()
		}

		return values[0].Bool(), values[1]
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func FindBy(source interface{}, properties map[string]interface{}) (interface{}, error) {
	if source == nil || properties == nil || len(properties) == 0 {
		return nil, nil
	}

	return Find(source, func (item, _ interface{}) (bool, error) {
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
func (this *Query) Find(predicate interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = Find(this.source, predicate)
	}
	return this
}

func (this *Query) FindBy(properties map[string]interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = FindBy(this.source, properties)
	}
	return this
}