package underscore

import (
	"testing"
)

func TestIndex(t *testing.T) {
	res, _ := Index([]string{ "a", "b" }, func (item, _ interface{}) (interface{}, error) {
		return item, nil
	})
	str, ok := res["a"].(string)
	if !(ok && str == "a") {
		t.Error("index error")
	}
}

func TestChain_Index(t *testing.T) {
	v, _ := Chain([]string{ "a", "b" }).Index(func (item, _ interface{}) (interface{}, error) {
		return item, nil
	}).Value()
	dict, ok := v.(map[interface{}]interface{})
	if !(ok && len(dict) == 2) {
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
	dict, err := IndexBy(arr, "Name")
	if !(err == nil && len(dict) == 2) {
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

	dict, ok := res.(map[interface{}]interface{})
	if !(ok && len(dict) == 2) {
		t.Error("wrong")
	}
}