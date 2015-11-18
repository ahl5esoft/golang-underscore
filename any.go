package underscore

import (
	"errors"
	"reflect"
)

func Any(source, predicate interface{}) (bool, error) {
	predicateRV := reflect.ValueOf(predicate)
	if predicateRV.Kind() != reflect.Func {
		return false, errors.New("underscore: Any's predicate is not func")
	}

	var ok bool
	err := each(source, func (args []reflect.Value) (bool, reflect.Value) {
		values := predicateRV.Call(args)
		ok = values[0].Bool()
		return ok, values[1]
	})
	if err != nil {
		return false, err
	}

	return ok, nil
}

func AnyBy(source interface{}, properties map[string]interface{}) (bool, error) {
	if source == nil || properties == nil || len(properties) == 0 {
		return false, nil
	}

	return Any(source, func (item, _ interface{}) (bool, error) {
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
func (this *Query) Any(predicate interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = Any(this.source, predicate)
	}
	return this
}

func (this *Query) AnyBy(properties map[string]interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = AnyBy(this.source, properties)
	}
	return this
}