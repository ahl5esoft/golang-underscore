package underscore

import (
	"reflect"
)

func Pluck(source interface{}, field string) ([]interface{}, error) {
	if source == nil {
		return EMPTY_ARRAY, nil
	}

	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			if sourceRV.Len() == 0 {
				return EMPTY_ARRAY, nil
			}

			results := make([]interface{}, sourceRV.Len())
			for i := 0; i < sourceRV.Len(); i++ {
				v, err := getFieldValue(sourceRV.Index(i).Interface(), field)
				if err != nil {
					return EMPTY_ARRAY, err
				}

				results[i] = v
			}
			return results, nil
		case reflect.Map:
			oldKeyRVs := sourceRV.MapKeys()
			if len(oldKeyRVs) == 0 {
				return EMPTY_ARRAY, nil
			}

			results := make([]interface{}, len(oldKeyRVs))
			for i := 0; i < len(oldKeyRVs); i++ {
				v, err := getFieldValue(sourceRV.MapIndex(oldKeyRVs[i]).Interface(), field)
				if err != nil {
					return EMPTY_ARRAY, err
				}

				results[i] = v
			}
			return results, nil
	}
	return EMPTY_ARRAY, nil
}

//chain
func (this *Query) Pluck(field string) Queryer {
	if this.err == nil {
		this.source, this.err = Pluck(this.source, field)
	}
	return this
}