package underscore

import (
	"reflect"
)

func First(source interface{}) (interface{}, error) {
	v, _ := Take(source, 1)
	rv := reflect.ValueOf(v)
	if v == nil || rv.Len() == 0 {
		return nil, nil
	}

	return rv.Index(0).Interface(), nil
}

func Take(source interface{}, count int) (interface{}, error) {
	var err error
	errRV := reflect.ValueOf(err)

	var arrRV reflect.Value
	each(source, func (args []reflect.Value) (bool, reflect.Value) {
		if !arrRV.IsValid() {
			arrRV = makeSliceRVWithElem(args[0].Type(), 0)
		}

		arrRV = reflect.Append(arrRV, args[0])
		return arrRV.Len() > count, errRV
	})
	if arrRV.IsValid() {
		return arrRV.Interface(), nil
	}

	return nil, nil
}

//# chain
func (this *Query) First() Queryer {
	if this.err == nil {
		this.source, this.err = First(this.source)
	}
	return this
}

func (this *Query) Take(count int) Queryer {
	if this.err == nil {
		this.source, this.err = Take(this.source, count)
	}
	return this
}