package underscore

import (
	"testing"
)

func Test_Sort(t *testing.T) {
	arr := []TestModel{
		{ID: 2, Name: "two"},
		{ID: 1, Name: "one"},
		{ID: 3, Name: "three"},
	}
	res := make([]TestModel, 0)
	Sort(arr, func(n TestModel, _ int) int {
		return n.ID
	}, &res)
	if len(res) != len(arr) {
		t.Fatal("wrong length")
	}

	if !(res[0].ID == 1 && res[1].ID == 2 && res[2].ID == 3) {
		t.Fatal("sort error")
	}
}

func Test_SortBy(t *testing.T) {
	arr := []TestModel{
		{ID: 2, Name: "two"},
		{ID: 1, Name: "one"},
		{ID: 3, Name: "three"},
	}
	res := make([]TestModel, 0)
	SortBy(arr, "id", &res)
	if len(res) != len(arr) {
		t.Fatal("wrong length")
	}

	if !(res[0].ID < res[1].ID && res[1].ID < res[2].ID) {
		t.Error("wrong result")
	}
}

func Test_Chain_Sort(t *testing.T) {
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

func Test_Chain_SortBy(t *testing.T) {
	arr := []TestModel{
		{ID: 3, Name: "three"},
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
	}
	res := make([]TestModel, 0)
	Chain(arr).SortBy("id").Value(&res)
	if !(len(res) == len(arr) && res[0].ID < res[1].ID && res[1].ID < res[2].ID) {
		t.Error(res)
	}
}
