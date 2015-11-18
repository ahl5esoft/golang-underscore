package underscore

import (
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{ 1, 2, 3, 5 }
	v, _ := Sort([]int{ 5, 3, 2, 1 }, func (n, _ int) (int, error) {
		return n, nil
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
	v, _ := Chain([]int{ 5, 3, 2, 1 }).Sort(func (n, _ int) (int, error) {
		return n, nil
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
	v, _ := SortBy(arr, "Id")
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
	v, _ := Chain(arr).SortBy("Id").Value()
	res, ok := v.([]TestModel)
	if !(ok && len(res) == len(arr)) {
		t.Error("wrong length")
		return
	}

	if !(res[0].Id < res[1].Id && res[1].Id < res[2].Id) {
		t.Error("wrong result")
	}
}