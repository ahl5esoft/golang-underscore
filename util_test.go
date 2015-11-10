package underscore

import (
	"reflect"
	"testing"
)

type TestModel struct {
	Id int
	Name string
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

	t.Log(
		reflect.SliceOf(
			reflect.TypeOf(1),
		),
	)
}