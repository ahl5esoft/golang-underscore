package underscore

import (
	"testing"
)

func TestGroup(t *testing.T) {
	v := Group([]int{ 1, 2, 3, 4, 5 }, func (n, _ int) string {
		if n % 2 == 0 {
			return "even"
		}
		return "odd"
	})
	dict, ok := v.(map[string][]int)
	if !(ok && len(dict["even"]) == 2) {
		t.Error("wrong")
	}
}

func TestChain_Group(t *testing.T) {
	v := Chain([]int{ 1, 2, 3, 4, 5 }).Group(func (n, _ int) string {
		if n % 2 == 0 {
			return "even"
		}
		return "odd"
	}).Value()
	dict, ok := v.(map[string][]int)
	if !(ok && len(dict["even"]) == 2) {
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
	v := GroupBy(arr, "name")
	dict, ok := v.(map[string][]TestModel)
	if !(ok && len(dict) == 2) {
		t.Error("wrong")
	}
}

func TestChain_GroupBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "a" },
		TestModel{ 2, "a" },
		TestModel{ 3, "b" },
		TestModel{ 4, "b" },
	}
	v := Chain(arr).GroupBy("Name").Value()
	dict, ok := v.(map[string][]TestModel)
	if !(ok && len(dict) == 2) {
		t.Error("wrong")
	}
}