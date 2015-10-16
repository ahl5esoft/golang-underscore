package underscore

import (
	"errors"
	"reflect"
)

func Select(source interface{}, predicate func(interface{}, interface{}) (bool, error)) ([]interface{}, error) {
	if predicate == nil {
		return EMPTY_ARRAY, errors.New("underscore: Select's predicate is nil")
	}

	if source == nil {
		return EMPTY_ARRAY, nil
	}

	results := []interface{}{}
	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			if sourceRV.Len() == 0 {
				return EMPTY_ARRAY, nil
			}

			for i := 0; i < sourceRV.Len(); i++ {
				item := sourceRV.Index(i).Interface()
				ok, err := predicate(item, i)
				if err != nil {
					return EMPTY_ARRAY, err
				}

				if ok {
					results = append(results, item)
				}
			}
		case reflect.Map:
			keyRVs := sourceRV.MapKeys()
			if len(keyRVs) == 0 {
				return EMPTY_ARRAY, nil
			}

			for i := 0; i < len(keyRVs); i++ {
				item := sourceRV.MapIndex(keyRVs[i]).Interface()
				ok, err := predicate(
					item,
					keyRVs[i].Interface(),
				)
				if err != nil {
					return EMPTY_ARRAY, err
				}

				if ok {
					results = append(results, item)
				}
			}
	}
	return results, nil
}

func SelectBy(source interface{}, properties map[string]interface{}) ([]interface{}, error) {
	if source == nil || properties == nil || len(properties) == 0 {
		return EMPTY_ARRAY, nil
	}

	return Select(source, func (item, _ interface{}) (bool, error) {
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
func (this *Query) Select(predicate func(interface{}, interface{}) (bool, error)) Queryer {
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