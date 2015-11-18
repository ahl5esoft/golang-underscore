package underscore

import (
	"errors"
	"reflect"
	"testing"
)

type TestModel struct {
	Id int
	Name string
}

func Test_each(t *testing.T) {
	indexes := []int{ 1, 3, 5, 7 }
	err := each(indexes, func (rvs []reflect.Value) (bool, reflect.Value) {
		i := rvs[1].Int()
		if rvs[0].Int() != int64(indexes[i]) {
			return false, reflect.ValueOf(
				errors.New("item"),
			)
		}
		return false, reflect.ValueOf(nil)
	})
	if err != nil {
		t.Error(err)
	}
}

func Test_getPropertyValue(t *testing.T) {
	entity := TestModel{ 1, "name" }
	_, err := getPropertyValue(entity, "Id")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = getPropertyValue(entity, "id")
	if err == nil {
		t.Error("fail")
	}
}