package underscore

import (
	"testing"
)

func TestWhere(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
		TestModel{ID: 4, Name: "three"},
	}
	res := make([]TestModel, 0)
	Where(arr, func(r TestModel, i int) bool {
		return r.ID%2 == 0
	}, &res)
	if !(len(res) == 2 && res[0].ID == 2 && res[1].ID == 4) {
		t.Error(res)
	}
}

func TestWhereBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "one"},
		TestModel{ID: 3, Name: "three"},
		TestModel{ID: 4, Name: "three"},
	}
	res := make([]TestModel, 0)
	WhereBy(arr, map[string]interface{}{
		"Name": "one",
	}, &res)
	if !(len(res) == 2 && res[0] == arr[0] && res[1] == arr[1]) {
		t.Error(res)
	}
}

func TestChain_Where(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "one"},
		TestModel{ID: 3, Name: "three"},
		TestModel{ID: 4, Name: "three"},
	}
	res := make([]TestModel, 0)
	Chain(arr).Where(func(r TestModel, i int) bool {
		return r.ID%2 == 0
	}).Value(&res)
	if !(len(res) == 2 && res[0].ID == 2 && res[1].ID == 4) {
		t.Error("wrong result")
	}
}

func TestChain_WhereBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "one"},
		TestModel{ID: 3, Name: "three"},
		TestModel{ID: 4, Name: "three"},
	}
	res := make([]TestModel, 0)
	Chain(arr).WhereBy(map[string]interface{}{
		"Name": "one",
	}).Value(&res)
	if !(len(res) == 2 && res[0] == arr[0] && res[1] == arr[1]) {
		t.Error("wrong result")
	}
}
