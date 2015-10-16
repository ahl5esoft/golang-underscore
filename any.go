package underscore

import (
	"errors"
	"reflect"
)

func Any(source interface{}, predicate func(interface{}, interface{}) (bool, error)) (bool, error) {
	if predicate == nil {
		return false, errors.New("underscore: Any's predicate is nil")
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
				ok, err := predicate(
					sourceRV.Index(i).Interface(),
					i,
				)
				if err == nil && ok {
					return true, nil
				}
			}
		case reflect.Map:
			keyRVs := sourceRV.MapKeys()
			if len(keyRVs) == 0 {
				return false, nil
			}

			for i := 0; i < len(keyRVs); i++ {
				ok, err := predicate(
					sourceRV.MapIndex(keyRVs[i]).Interface(),
					keyRVs[i].Interface(),
				)
				if err == nil && ok {
					return true, nil
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
		return All(properties, func (pv, pn interface{}) (bool, error) {
			value, err := getFieldValue(item, pn.(string))
			if err != nil {
				return false, err
			}

			return value == pv, nil
		})
	})
}

//# chain
func (this *Query) Any(predicate func(interface{}, interface{}) (bool, error)) Queryer {
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