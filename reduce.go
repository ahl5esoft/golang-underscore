package underscore

import (
	"errors"
	"reflect"
)

func Reduce(source, iterator, memo interface{}) (interface{}, error) {
	iteratorRV := reflect.ValueOf(iterator)
	if iteratorRV.Kind() != reflect.Func {
		return nil, errors.New("underscore: Reduce's iterator is not func")
	}

	memoRV := reflect.ValueOf(memo)
	err := each(source, func (args []reflect.Value) (bool, reflect.Value) {
		values := iteratorRV.Call(
			append([]reflect.Value{
				memoRV,
			}, args...),
		)
		if !isErrorRVValid(values[1]) {
			memoRV = values[0]
		}

		return false, values[1]
	})
	if err == nil  {
		return memoRV.Interface(), nil
	}

	return nil, err
}

//Chain
func (this *Query) Reduce(iterator, memo interface{}) Queryer {
	if this.err == nil {
		this.source, this.err = Reduce(this.source, iterator, memo)
	}
	return this
}