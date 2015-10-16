package underscore

import (
	"errors"
	"reflect"
)

func All(source interface{}, predicate func(interface{}, interface{}) (bool, error)) (bool, error) {
	if predicate == nil {
		return false, errors.New("underscore: All's predicate is nil")
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
				ok, err := predicate(
					sourceRV.Index(i).Interface(),
					i,
				)
				if !(err == nil && ok) {
					return ok, err
				}
			}
		case reflect.Map:
			keyRVs := sourceRV.MapKeys()
			if len(keyRVs) == 0 {
				return true, nil
			}

			for i := 0; i < len(keyRVs); i++ {
				ok, err := predicate(
					sourceRV.MapIndex(keyRVs[i]).Interface(),
					keyRVs[i].Interface(),
				)
				if !(err == nil && ok) {
					return ok, err
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
		return All(properties, func (pv, pn interface{}) (bool, error) {
			value, err := getPropertyValue(item, pn.(string))
			if err != nil {
				return false, err
			}

			return value == pv, nil
		})
	})
}

//# chain
func (this *Query) All(predicate func(interface{}, interface{}) (bool, error)) Queryer {
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