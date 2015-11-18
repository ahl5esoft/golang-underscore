package underscore

import (
	"errors"
	"reflect"
	"sort"
)

type sortQuery struct {
	keysRV reflect.Value
	valuesRV reflect.Value
	compareRV reflect.Value
}

func (this sortQuery) Len() int {
	if this.keysRV.IsValid() {
		return this.keysRV.Len()
	}

	return 0;
}

func (this sortQuery) Swap(i, j int) {
	temp := this.keysRV.Index(i).Interface()
	this.keysRV.Index(i).Set(
		this.keysRV.Index(j),
	)
	this.keysRV.Index(j).Set(
		reflect.ValueOf(temp),
	)

	temp = this.valuesRV.Index(i).Interface()
	this.valuesRV.Index(i).Set(
		this.valuesRV.Index(j),
	)
	this.valuesRV.Index(j).Set(
		reflect.ValueOf(temp),
	)
}

func (this sortQuery) Less(i, j int) bool {
	thisRV := this.keysRV.Index(i)
	thatRV := this.keysRV.Index(j)
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

func Sort(source, selector interface{}) (interface{}, error) {
	selectorRV := reflect.ValueOf(selector)
	if selectorRV.Kind() != reflect.Func {
		return nil, errors.New("underscore: Sort's selector is not func")
	}

	qs := sortQuery{}
	err := each(source, func (args []reflect.Value) (bool, reflect.Value) {
		if qs.Len() == 0 {
			qs.valuesRV = makeSliceRVWithElem(
				args[0].Type(),
				0,
			)
			qs.keysRV = makeSliceRVWithElem(
				selectorRV.Type().Out(0),
				0,
			)
		}

		values := selectorRV.Call(args)
		if !isErrorRVValid(values[1]) {
			qs.valuesRV = reflect.Append(qs.valuesRV, args[0])
			qs.keysRV = reflect.Append(qs.keysRV, values[0])
		}
		
		return false, values[1]
	})
	if err == nil && qs.Len() > 0 {
		sort.Sort(qs)
		return qs.valuesRV.Interface(), nil
	}

	return nil, err
}

func SortBy(source interface{}, property string) (interface{}, error) {
	qs := sortQuery{}
	err := each(source, func (args []reflect.Value) (bool, reflect.Value) {
		pRV, err := getPropertyRV(args[0], property)
		if err == nil {
			if qs.Len() == 0 {
				qs.valuesRV = makeSliceRVWithElem(
					args[0].Type(),
					0,
				)
				qs.keysRV = makeSliceRVWithElem(
					pRV.Type(),
					0,
				)
			}

			qs.valuesRV = reflect.Append(qs.valuesRV, args[0])
			qs.keysRV = reflect.Append(qs.keysRV, pRV)		
		}

		return false, reflect.ValueOf(err)
	})
	if err == nil && qs.Len() > 0 {
		sort.Sort(qs)
		return qs.valuesRV.Interface(), nil
	}

	return nil, err
}

//chain
func (this *Query) Sort(selector interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = Sort(this.source, selector)
	}
	return this
}

func (this *Query) SortBy(property string) Queryer {
	if this.err == nil {
		this.source, this.err = SortBy(this.source, property)
	}
	return this
}