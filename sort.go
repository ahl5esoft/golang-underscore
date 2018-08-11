package underscore

import (
	"reflect"
	"sort"
)

type sortQuery struct {
	keysRV    reflect.Value
	valuesRV  reflect.Value
	compareRV reflect.Value
}

func (q sortQuery) Len() int {
	if q.keysRV.IsValid() {
		return q.keysRV.Len()
	}

	return 0
}

func (q sortQuery) Swap(i, j int) {
	temp := q.keysRV.Index(i).Interface()
	q.keysRV.Index(i).Set(
		q.keysRV.Index(j),
	)
	q.keysRV.Index(j).Set(
		reflect.ValueOf(temp),
	)

	temp = q.valuesRV.Index(i).Interface()
	q.valuesRV.Index(i).Set(
		q.valuesRV.Index(j),
	)
	q.valuesRV.Index(j).Set(
		reflect.ValueOf(temp),
	)
}

func (q sortQuery) Less(i, j int) bool {
	thisRV := q.keysRV.Index(i)
	thatRV := q.keysRV.Index(j)
	switch thisRV.Kind() {
	case reflect.Float32, reflect.Float64:
		return thisRV.Float() < thatRV.Float()
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return thisRV.Int() < thatRV.Int()
	case reflect.String:
		return thisRV.String() < thatRV.String()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return thisRV.Uint() < thatRV.Uint()
	default:
		return false
	}
}

// Sort is 排序
func Sort(source, selector interface{}) interface{} {
	qs := sortQuery{}
	each(source, selector, func(sortRV, valueRV, _ reflect.Value) bool {
		if qs.Len() == 0 {
			keysRT := reflect.SliceOf(sortRV.Type())
			qs.keysRV = reflect.MakeSlice(keysRT, 0, 0)

			valuesRT := reflect.SliceOf(valueRV.Type())
			qs.valuesRV = reflect.MakeSlice(valuesRT, 0, 0)
		}

		qs.keysRV = reflect.Append(qs.keysRV, sortRV)
		qs.valuesRV = reflect.Append(qs.valuesRV, valueRV)
		return false
	})
	if qs.Len() > 0 {
		sort.Sort(qs)
		return qs.valuesRV.Interface()
	}

	return nil
}

// SortBy is 根据属性排序
func SortBy(source interface{}, property string) interface{} {
	getPropertyRV := PropertyRV(property)
	return Sort(source, func(value, _ interface{}) Facade {
		rv, _ := getPropertyRV(value)
		return Facade{rv}
	})
}

// Sort is Queryer's method
func (q *Query) Sort(selector interface{}) Queryer {
	q.source = Sort(q.source, selector)
	return q
}

// SortBy is Queryer's method
func (q *Query) SortBy(property string) Queryer {
	q.source = SortBy(q.source, property)
	return q
}
