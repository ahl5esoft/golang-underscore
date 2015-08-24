package underscore

import (
	"errors"
	"reflect"
	"sort"
)

type sortQuery struct {
	keys []interface{}
	values []interface{}
	compare func(thisValue, thisKey, thatValue, thatKey interface{}) bool
}

func (this sortQuery) Len() int {
	return len(this.values);
}

func (this sortQuery) Swap(i, j int) {
	this.keys[i], this.keys[j] = this.keys[j], this.keys[i]
	this.values[i], this.values[j] = this.values[j], this.values[i]
}

func (this sortQuery) Less(i, j int) bool {
	return this.compare(this.values[i], this.keys[i], this.values[j], this.keys[j])
}

func Sort(source interface{}, compare func(thisValue, thisKey, thatValue, thatKey interface{}) bool) ([]interface{}, error) {
	if compare == nil {
		return EMPTY_ARRAY, errors.New("underscore: Sort's compare is nil")
	}

	sourceRV := reflect.ValueOf(source)
	if sourceRV.Kind() == reflect.Array || sourceRV.Kind() == reflect.Slice || sourceRV.Kind() == reflect.Map {
		qs := sortQuery{}
		qs.compare = compare

		if sourceRV.Kind() == reflect.Map {
			oldKeyRVs := sourceRV.MapKeys()
			if len(oldKeyRVs) == 0 {
				return EMPTY_ARRAY, nil
			}

			qs.keys = make([]interface{}, len(oldKeyRVs))
			qs.values = make([]interface{}, len(oldKeyRVs))
			for i := 0; i < len(oldKeyRVs); i++ {
				qs.keys[i] = oldKeyRVs[i].Interface()
				qs.values[i] = sourceRV.MapIndex(oldKeyRVs[i]).Interface()
			}
		} else {
			if sourceRV.Len() == 0 {
				return EMPTY_ARRAY, nil
			}

			qs.keys = make([]interface{}, sourceRV.Len())
			qs.values = make([]interface{}, sourceRV.Len())
			for i := 0; i < sourceRV.Len(); i++ {
				qs.keys[i] = i
				qs.values[i] = sourceRV.Index(i).Interface()
			}
		}

		sort.Sort(qs)
		return qs.values, nil
	}
	return EMPTY_ARRAY, nil
}

func SortBy(source interface{}, field string) ([]interface{}, error) {
	sourceRV := reflect.ValueOf(source)
	if sourceRV.Kind() == reflect.Array || sourceRV.Kind() == reflect.Slice || sourceRV.Kind() == reflect.Map {
		var fieldKind reflect.Kind
		qs := sortQuery{}
		qs.compare = func (_, thisKey, _, thatKey interface{}) bool {
			switch fieldKind {
				case reflect.Float32:
					return thisKey.(float32) < thatKey.(float32)
				case reflect.Float64:
					return thisKey.(float64) < thatKey.(float64)
				case reflect.Int:
					return thisKey.(int) < thatKey.(int)
				case reflect.Int16:
					return thisKey.(int16) < thatKey.(int16)
				case reflect.Int32:
					return thisKey.(int) < thatKey.(int)
				case reflect.Int64:
					return thisKey.(int64) < thatKey.(int64)
				case reflect.String:
					return thisKey.(string) < thatKey.(string)
				default:
					return false
			}
		}

		if sourceRV.Kind() == reflect.Map {
			oldKeyRVs := sourceRV.MapKeys()
			if len(oldKeyRVs) == 0 {
				return EMPTY_ARRAY, nil
			}

			qs.keys = make([]interface{}, len(oldKeyRVs))
			qs.values = make([]interface{}, len(oldKeyRVs))
			for i := 0; i < len(oldKeyRVs); i++ {
				qs.values[i] = sourceRV.MapIndex(oldKeyRVs[i]).Interface()
				value, err := getFieldValue(qs.values[i], field)
				if err != nil {
					return EMPTY_ARRAY, err
				}

				if i == 0 {
					fieldKind = reflect.ValueOf(qs.values[i]).FieldByName(field).Kind()
				}
				qs.keys[i] = value
			}
		} else {
			if sourceRV.Len() == 0 {
				return EMPTY_ARRAY, nil
			}

			qs.keys = make([]interface{}, sourceRV.Len())
			qs.values = make([]interface{}, sourceRV.Len())
			for i := 0; i < sourceRV.Len(); i++ {
				qs.values[i] = sourceRV.Index(i).Interface()
				value, err := getFieldValue(qs.values[i], field)
				if err != nil {
					return EMPTY_ARRAY, nil
				}

				if i == 0 {
					fieldKind = reflect.ValueOf(qs.values[i]).FieldByName(field).Kind()
				}
				qs.keys[i] = value
			}
		}

		sort.Sort(qs)
		return qs.values, nil
	}
	return EMPTY_ARRAY, nil
}

//chain
func (this *Query) Sort(compare func(thisValue, thisKey, thatValue, thatKey interface{}) bool) Queryer {
	if this.err == nil {
		this.source, this.err = Sort(this.source, compare)
	}
	return this
}

func (this *Query) SortBy(field string) Queryer {
	if this.err == nil {
		this.source, this.err = SortBy(this.source, field)
	}
	return this
}