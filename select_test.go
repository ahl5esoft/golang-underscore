package underscore

import (
	"testing"
)

func TestSelect(t *testing.T) {
	arr := []int{ 1, 2, 3, 4 }
	v := Select(arr, func (n, i int) bool {
		return n % 2 == 0
	})
	res, ok := v.([]int)
	if !(ok && len(res) == 2) {
		t.Error("wrong length")
		return
	}

	if !(res[0] == 2 && res[1] == 4) {
		t.Error("wrong result")
	}
}

func TestChain_Select(t *testing.T) {
	arr := []int{ 1, 2, 3, 4 }
	v := Chain(arr).Select(func (n, i int) bool {
		return n % 2 == 0
	}).Value()
	res, ok := v.([]int)
	if !(ok && len(res) == 2) {
		t.Error("wrong length")
		return
	}

	if !(res[0] == 2 && res[1] == 4) {
		t.Error("wrong result")
	}
}

func TestSelectBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	v := SelectBy(arr, map[string]interface{}{
		"Id": 1,
	})
	res, ok := v.([]TestModel)
	if !(ok && len(res) == 1 && res[0] == arr[0]) {
		t.Error("wrong result")
	}
}

func TestChain_SelectBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	v := Chain(arr).SelectBy(map[string]interface{}{
		"Id": 1,
	}).Value()
	res, ok := v.([]TestModel)
	if !(ok && len(res) == 1 && res[0] == arr[0]) {
		t.Error("wrong result")
	}
}