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

	if source == nil  {
		return false, nil
	}

	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			if sourceRV.Len() == 0 {
				return false, nil
			}

			for i := 0; i < sourceRV.Len(); i++ {
				values := predicateRV.Call(
					[]reflect.Value{
						sourceRV.Index(i),
						reflect.ValueOf(i),
					},
				)
				if values[0].Bool() && values[1].IsNil() {
					return true, nil
				} else if !values[1].IsNil() {
					return false, values[1].Interface().(error)
				}
			}
		case reflect.Map:
			keyRVs := sourceRV.MapKeys()
			if len(keyRVs) == 0 {
				return false, nil
			}

			for i := 0; i < len(keyRVs); i++ {
				values := predicateRV.Call(
					[]reflect.Value{
						sourceRV.MapIndex(keyRVs[i]),
						keyRVs[i],
					},
				)
				if values[0].Bool() && values[1].IsNil() {
					return true, nil
				} else if !values[1].IsNil() {
					return false, values[1].Interface().(error)
				}
			}
	}
	return false, nil
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