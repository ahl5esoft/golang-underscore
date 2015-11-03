package underscore

import (
	"errors"
	"reflect"
)

func Find(source interface{}, predicate func(interface{}, interface{}) (bool, error)) (interface{}, error) {
	if predicate == nil {
		return nil, errors.New("underscore: Find's predicate is nil")
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
				item := sourceRV.Index(i).Interface()
				ok, err := predicate(item, i)
				if err != nil {
					return nil, err
				}

				if ok {
					return item, nil
				}
			}
		case reflect.Map:
			keyRVs := sourceRV.MapKeys()
			if len(keyRVs) == 0 {
				return nil, nil
			}

			for i := 0; i < len(keyRVs); i++ {
				item := sourceRV.MapIndex(keyRVs[i]).Interface()
				ok, err := predicate(
					item,
					keyRVs[i].Interface(),
				)
				if err != nil {
					return nil, err
				}

				if ok {
					return item, nil
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
func (this *Query) Find(predicate func(interface{}, interface{}) (bool, error)) Queryer {
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