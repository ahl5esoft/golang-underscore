package underscore

import (
	"testing"
)

func Test_Index(t *testing.T) {
	res := make(map[string]string)
	Index([]string{"a", "b"}, func(r string, _ int) string {
		return r
	}, &res)
	if res["a"] != "a" {
		t.Error(res)
	}
}

func Test_IndexBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "a"},
		TestModel{ID: 2, Name: "a"},
		TestModel{ID: 3, Name: "b"},
		TestModel{ID: 4, Name: "b"},
	}
	res := make(map[string]TestModel)
	IndexBy(arr, "Name", &res)
	if len(res) != 2 {
		t.Error(res)
	}
}

func Test_Chain_Index(t *testing.T) {
	res := make(map[string]string)
	Chain([]string{"a", "b"}).Index(func(item string, _ int) string {
		return item
	}).Value(&res)
	if res["a"] != "a" {
		t.Error(res)
	}
}

func Test_Chain_IndexBy(t *testing.T) {
	res := make(map[string]TestModel)
	Chain([]TestModel{
		TestModel{ID: 1, Name: "a"},
		TestModel{ID: 2, Name: "a"},
		TestModel{ID: 3, Name: "b"},
		TestModel{ID: 4, Name: "b"},
	}).IndexBy("Name").Value(&res)
	if len(res) != 2 {
		t.Error("wrong")
	}
}
