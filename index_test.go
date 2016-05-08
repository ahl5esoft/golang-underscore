package underscore

import (
	"testing"
)

func TestIndex(t *testing.T) {
	v := Index([]string{ "a", "b" }, func (item string, _ int) string {
		return item
	})
	res, ok := v.(map[string]string)
	if !(ok && res["a"] == "a") {
		t.Error("wrong")
	}
}

func TestChain_Index(t *testing.T) {
	v := Chain([]string{ "a", "b" }).Index(func (item string, _ int) string {
		return item
	}).Value()
	res, ok := v.(map[string]string)
	if !(ok && res["a"] == "a") {
		t.Error("wrong")
	}
}

func TestIndexBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "a" },
		TestModel{ 2, "a" },
		TestModel{ 3, "b" },
		TestModel{ 4, "b" },
	}
	res := IndexBy(arr, "Name")
	dict, ok := res.(map[string]TestModel)
	if !(ok && len(dict) == 2) {
		t.Error("wrong")
	}
}

func TestChain_IndexBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "a" },
		TestModel{ 2, "a" },
		TestModel{ 3, "b" },
		TestModel{ 4, "b" },
	}
	res := Chain(arr).IndexBy("Name").Value()
	dict, ok := res.(map[string]TestModel)
	if !(ok && len(dict) == 2) {
		t.Error("wrong")
	}
}