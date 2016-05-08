package underscore

import (
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{ 1, 2, 3, 5 }
	v := Sort([]int{ 5, 3, 2, 1 }, func (n, _ int) int {
		return n
	})
	res, ok := v.([]int)
	if !(ok && len(res) == len(arr)) {
		t.Error("wrong length")
		return
	}

	for i, n := range arr {
		if res[i] != n {
			t.Error("wrong result")
			break
		}
	}
}

func TestChain_Sort(t *testing.T) {
	arr := []int{ 1, 2, 3, 5 }
	v := Chain([]int{ 5, 3, 2, 1 }).Sort(func (n, _ int) int {
		return n
	}).Value()

	res, ok := v.([]int)
	if !(ok && len(res) == len(arr)) {
		t.Error("wrong length")
	}

	for i, n := range arr {
		if res[i] != n {
			t.Error("wrong result")
			break
		}
	}
}

func TestSortBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 3, "three" },
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
	}
	v := SortBy(arr, "id")
	res, ok := v.([]TestModel)
	if !(ok && len(res) == len(arr)) {
		t.Error("wrong length")
		return
	}

	if !(res[0].Id < res[1].Id && res[1].Id < res[2].Id) {
		t.Error("wrong result")
	}
}

func TestChain_SortBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 3, "three" },
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
	}
	v := Chain(arr).SortBy("id").Value()
	res, ok := v.([]TestModel)
	if !(ok && len(res) == len(arr)) {
		t.Error("wrong length")
		return
	}

	if !(res[0].Id < res[1].Id && res[1].Id < res[2].Id) {
		t.Error("wrong result")
	}
}