package underscore

import (
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	arr := []string{ "11", "12", "13" }
	v := Map(arr, func (s string, _ int) int {
		n, _ := strconv.Atoi(s)
		return n
	})
	res, ok := v.([]int)
	if !(ok && len(res) == len(arr)) {
		t.Error("wrong")
	}
}

func TestChain_Map(t *testing.T) {
	arr := []string{ "a", "b", "c" }
	v := Chain(arr).Map(func (item, _ interface{}) string {
		return item.(string) + "-"
	}).Value()
	res, ok := v.([]string)
	if !(ok && len(res) == len(arr) && res[0] == "a-") {
		t.Error("wrong")
	}
}

func TestMapBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "a" },
		TestModel{ 2, "a" },
		TestModel{ 3, "b" },
		TestModel{ 4, "b" },
	}
	v := MapBy(arr, "name")
	res, ok := v.([]string)
	if !(ok && len(res) == 4) {
		t.Error("wrong")
	}
}

func TestChain_MapBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "a" },
		TestModel{ 2, "a" },
		TestModel{ 3, "b" },
		TestModel{ 4, "b" },
	}
	v := Chain(arr).MapBy("id").Value()
	res, ok := v.([]int)
	if !(ok && len(res) == 4) {
		t.Error("wrong")
	}
}