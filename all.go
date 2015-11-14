package underscore

import (
	"errors"
	"reflect"
)

func All(source, predicate interface{}) (bool, error) {
	predicateRV := reflect.ValueOf(predicate)
	if predicateRV.Kind() != reflect.Func {
		return false, errors.New("underscore: All's predicate is not func")
	}

	if source == nil  {
		return true, nil
	}

	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			if sourceRV.Len() == 0 {
				return true, nil
			}

			for i := 0; i < sourceRV.Len(); i++ {
				values := predicateRV.Call(
					[]reflect.Value{
						sourceRV.Index(i),
						reflect.ValueOf(i),
					},
				)
				if !values[1].IsNil() {
					return false, values[1].Interface().(error)
				}
				
				if !values[0].Bool() {
					return false, nil
				}
			}
		case reflect.Map:
			keyRVs := sourceRV.MapKeys()
			if len(keyRVs) == 0 {
				return true, nil
			}

			for i := 0; i < len(keyRVs); i++ {
				values := predicateRV.Call(
					[]reflect.Value{
						sourceRV.MapIndex(keyRVs[i]),
						keyRVs[i],
					},
				)
				if !values[1].IsNil() {
					return false, values[1].Interface().(error)
				}
				
				if !values[0].Bool() {
					return false, nil
				}
			}
	}
	return true, nil
}

func AllBy(source interface{}, properties map[string]interface{}) (bool, error) {
	if source == nil || properties == nil || len(properties) == 0 {
		return true, nil
	}

	return All(source, func (item, _ interface{}) (bool, error) {
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
func (this *Query) All(predicate interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = All(this.source, predicate)
	}
	return this
}

func (this *Query) AllBy(properties map[string]interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = AllBy(this.source, properties)
	}
	return this
}