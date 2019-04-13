package underscore

import (
	"testing"
)

func Test_Sort(t *testing.T) {
	arr := []testModel{
		{ID: 2, Name: "two"},
		{ID: 1, Name: "one"},
		{ID: 3, Name: "three"},
	}
	res := Sort(arr, func(n testModel, _ int) int {
		return n.ID
	}).([]testModel)
	if len(res) != len(arr) {
		t.Fatal("wrong length")
	}

	if !(res[0].ID == 1 && res[1].ID == 2 && res[2].ID == 3) {
		t.Fatal("sort error")
	}
}

func Test_SortBy(t *testing.T) {
	arr := []testModel{
		{ID: 2, Name: "two"},
		{ID: 1, Name: "one"},
		{ID: 3, Name: "three"},
	}
	res := SortBy(arr, "id").([]testModel)
	if len(res) != len(arr) {
		t.Fatal("wrong length")
	}

	if !(res[0].ID < res[1].ID && res[1].ID < res[2].ID) {
		t.Error("wrong result")
	}
}

func Test_Chain_Sort(t *testing.T) {
	arr := []testModel{
		{ID: 2, Name: "two"},
		{ID: 1, Name: "one"},
		{ID: 3, Name: "three"},
	}
	var res []testModel
	Chain(arr).Sort(func(n testModel, _ int) int {
		return n.ID
	}).Value(&res)
	if len(res) != len(arr) {
		t.Fatal("wrong length")
	}

	if !(res[0].ID == 1 && res[1].ID == 2 && res[2].ID == 3) {
		t.Fatal("sort error")
	}
}

func Test_Chain_SortBy(t *testing.T) {
	arr := []testModel{
		{ID: 3, Name: "three"},
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
	}
	res := make([]testModel, 0)
	Chain(arr).SortBy("id").Value(&res)
	if !(len(res) == len(arr) && res[0].ID < res[1].ID && res[1].ID < res[2].ID) {
		t.Error(res)
	}
}
