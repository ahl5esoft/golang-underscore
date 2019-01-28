package underscore

import (
	"testing"
)

func TestSort(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 3, Name: "three"},
	}
	v := Sort(arr, func(n TestModel, _ int) int {
		return -n.ID
	})
	res, ok := v.([]TestModel)
	if !(ok && len(res) == len(arr)) {
		t.Fatal("wrong length")
	}

	if !(res[0].ID == 3 && res[1].ID == 2 && res[2].ID == 1) {
		t.Fatal("sort error")
	}
}

func TestSortBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 3, Name: "three"},
	}
	v := SortBy(arr, "id")
	res, ok := v.([]TestModel)
	if !(ok && len(res) == len(arr)) {
		t.Error("wrong length")
		return
	}

	if !(res[0].ID < res[1].ID && res[1].ID < res[2].ID) {
		t.Error("wrong result")
	}
}

func TestChain_Sort(t *testing.T) {
	arr := []int{1, 2, 3, 5}
	res := make([]int, 0)
	Chain([]int{5, 3, 2, 1}).Sort(func(n, _ int) int {
		return n
	}).Value(&res)
	if len(res) != len(arr) {
		t.Fatal("wrong length")
	}

	for i, n := range arr {
		if res[i] != n {
			t.Fatal("wrong result")
		}
	}
}

func TestChain_SortBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 3, Name: "three"},
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
	}
	res := make([]TestModel, 0)
	Chain(arr).SortBy("id").Value(&res)
	if !(len(res) == len(arr) && res[0].ID < res[1].ID && res[1].ID < res[2].ID) {
		t.Error(res)
	}
}
