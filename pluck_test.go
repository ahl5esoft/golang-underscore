package underscore

import (
	"testing"
)

func TestPluck(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	res, err := Pluck(arr, "Name")
	if err != nil {
		t.Error(err)
	}

	if len(res) != len(arr) {
		t.Error("wrong length")
	}

	for i := 0; i < 3; i++ {
		if res[i].(string) != arr[i].Name {
			t.Error("wrong result")
		}
	}
}

func TestChain_Pluck(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	v, err := Chain(arr).Pluck("Name").Value()
	if err != nil {
		t.Error(err)
	}

	res, ok := v.([]interface{})
	if !(ok && len(res) == len(arr)) {
		t.Error("wrong length")
	}

	for i := 0; i < 3; i++ {
		if res[i].(string) != arr[i].Name {
			t.Error("wrong result")
		}
	}
}