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

func TestParseJson(t *testing.T) {
	str := `["a","b"]`
	var arr []string
	err := ParseJson(str, &arr)
	if !(err == nil && len(arr) == 2) {
		t.Error("wrong")
	}
}

func TestToJson(t *testing.T) {
	b := true
	v, _ := ToJson(b)
	if v != "true" {
		t.Error("bool fail")
		return
	}

	str := "a"
	v, _ = ToJson(str)
	if v != str {
		t.Error("string fail")
		return
	}

	v, _ = ToJson(1)
	if v != "1" {
		t.Error("int fail")
		return
	}

	arr := []int{ 1, 2, 3 }
	v, _ = ToJson(arr)
	if v != "[1,2,3]" {
		t.Error("array fail")
		return
	}

	obj := TestModel{ 1, "name" }
	v, _ = ToJson(obj)
	if v != `{"Id":1,"Name":"name"}` {
		t.Error("obj fail")
		return
	}
}