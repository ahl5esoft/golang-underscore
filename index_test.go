package underscore

import (
	"testing"
)

func TestIndex(t *testing.T) {
	v := Index([]string{"a", "b"}, func(item string, _ int) string {
		return item
	})
	res, ok := v.(map[string]string)
	if !(ok && res["a"] == "a") {
		t.Error("wrong")
	}
}

func TestIndexBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "a"},
		TestModel{ID: 2, Name: "a"},
		TestModel{ID: 3, Name: "b"},
		TestModel{ID: 4, Name: "b"},
	}
	res := IndexBy(arr, "Name")
	dict, ok := res.(map[string]TestModel)
	if !(ok && len(dict) == 2) {
		t.Error("wrong")
	}
}

func TestChain_Index(t *testing.T) {
	res := make(map[string]string)
	Chain([]string{"a", "b"}).Index(func(item string, _ int) string {
		return item
	}).Value(&res)
	if res["a"] != "a" {
		t.Error("wrong")
	}
}

func TestChain_IndexBy(t *testing.T) {
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
