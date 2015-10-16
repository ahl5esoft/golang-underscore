package underscore

import (
	"testing"
)

func TestGroup(t *testing.T) {
	dict, _ := Group([]int{ 1, 2, 3, 4, 5 }, func (item, _ interface{}) (interface{}, error) {
		if item.(int) % 2 == 0 {
			return "even", nil
		}
		return "odd", nil
	})
	group, ok := dict["even"]
	if !(ok && len(group) == 2) {
		t.Error("wrong")
	}
}

func TestChain_Group(t *testing.T) {
	v, _ := Chain([]int{ 1, 2, 3, 4, 5 }).Group(func (item, _ interface{}) (interface{}, error) {
		if item.(int) % 2 == 0 {
			return "even", nil
		}
		return "odd", nil
	}).Value()
	dict, ok := v.(map[interface{}][]interface{})
	if !(ok && len(dict) == 2) {
		t.Error("wrong")
	}
}

func TestGroupBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "a" },
		TestModel{ 2, "a" },
		TestModel{ 3, "b" },
		TestModel{ 4, "b" },
	}
	dict, err := GroupBy(arr, "Name")
	if !(err == nil && len(dict) == 2) {
		t.Error(err)
	}
}

func TestChain_GroupBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "a" },
		TestModel{ 2, "a" },
		TestModel{ 3, "b" },
		TestModel{ 4, "b" },
	}
	v, err := Chain(arr).GroupBy("Name").Value()
	if err != nil {
		t.Error("wrong")
	}

	dict, ok := v.(map[interface{}][]interface{})
	if !(ok && len(dict) == 2) {
		t.Error(err)
	}
}