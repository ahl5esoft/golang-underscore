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

	if source == nil {
		return nil, nil
	}

	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			if sourceRV.Len() == 0 {
				return nil, nil
			}

			for i := 0; i < sourceRV.Len(); i++ {
				values := predicateRV.Call(
					[]reflect.Value{
						sourceRV.Index(i),
						reflect.ValueOf(i),
					},
				)
				if values[0].Bool() && values[1].IsNil() {
					return sourceRV.Index(i).Interface(), nil
				} else if !values[1].IsNil() {
					return nil, values[0].Interface().(error)
				}
			}
		case reflect.Map:
			keyRVs := sourceRV.MapKeys()
			if len(keyRVs) == 0 {
				return nil, nil
			}

			for i := 0; i < len(keyRVs); i++ {
				values := predicateRV.Call(
					[]reflect.Value{
						sourceRV.MapIndex(keyRVs[i]),
						reflect.ValueOf(i),
					},
				)
				if values[0].Bool() && values[1].IsNil() {
					return sourceRV.MapIndex(keyRVs[i]).Interface(), nil
				} else if !values[1].IsNil() {
					return nil, values[0].Interface().(error)
				}
			}
	}
	return nil, nil
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