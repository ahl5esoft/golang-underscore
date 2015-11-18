package underscore

import (
	"testing"
)

func TestIndex(t *testing.T) {
	v, _ := Index([]string{ "a", "b" }, func (item string, _ int) (string, error) {
		return item, nil
	})
	res, ok := v.(map[string]string)
	if !(ok && res["a"] == "a") {
		t.Error("index error")
	}
}

func TestChain_Index(t *testing.T) {
	v, _ := Chain([]string{ "a", "b" }).Index(func (item string, _ int) (string, error) {
		return item, nil
	}).Value()
	res, ok := v.(map[string]string)
	if !(ok && res["a"] == "a") {
		t.Error("index error")
	}
}

func TestIndexBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "a" },
		TestModel{ 2, "a" },
		TestModel{ 3, "b" },
		TestModel{ 4, "b" },
	}
	res, err := IndexBy(arr, "Name")
	if err != nil {
		t.Error(err)
	}

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
	res, err := Chain(arr).IndexBy("Name").Value()
	if err != nil {
		t.Error(err)
	}

	dict, ok := res.(map[string]TestModel)
	if !(ok && len(dict) == 2) {
		t.Error("wrong")
	}
}