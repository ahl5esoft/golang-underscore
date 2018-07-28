package underscore

import (
	"testing"
)

func TestSelect(t *testing.T) {
	arr := []TestModel{
		TestModel{1, "one"},
		TestModel{2, "two"},
		TestModel{3, "three"},
		TestModel{4, "three"},
	}
	v := Where(arr, func(r TestModel, i int) bool {
		return r.Id%2 == 0
	})
	res, ok := v.([]TestModel)
	if !(ok && len(res) == 2) {
		t.Error("wrong length")
		return
	}

	if !(res[0].Id == 2 && res[1].Id == 4) {
		t.Error("wrong result")
	}
}

func TestChain_Select(t *testing.T) {
	arr := []TestModel{
		TestModel{1, "one"},
		TestModel{2, "two"},
		TestModel{3, "three"},
		TestModel{4, "three"},
	}
	v := Chain(arr).Where(func(r TestModel, i int) bool {
		return r.Id%2 == 0
	}).Value()
	res, ok := v.([]TestModel)
	if !(ok && len(res) == 2) {
		t.Error("wrong length")
		return
	}

	if !(res[0].Id == 2 && res[1].Id == 4) {
		t.Error("wrong result")
	}
}

func TestSelectBy(t *testing.T) {
	arr := []TestModel{
		TestModel{1, "one"},
		TestModel{2, "one"},
		TestModel{3, "three"},
		TestModel{4, "three"},
	}
	v := WhereBy(arr, map[string]interface{}{
		"Name": "one",
	})
	res, ok := v.([]TestModel)
	if !(ok && len(res) == 2 && res[0] == arr[0] && res[1] == arr[1]) {
		t.Error("wrong result")
	}
}

func TestChain_SelectBy(t *testing.T) {
	arr := []TestModel{
		TestModel{1, "one"},
		TestModel{2, "one"},
		TestModel{3, "three"},
		TestModel{4, "three"},
	}
	v := Chain(arr).WhereBy(map[string]interface{}{
		"Name": "one",
	}).Value()
	res, ok := v.([]TestModel)
	if !(ok && len(res) == 2 && res[0] == arr[0] && res[1] == arr[1]) {
		t.Error("wrong result")
	}
}
