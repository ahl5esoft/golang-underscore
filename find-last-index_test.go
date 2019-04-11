package underscore

import (
	"testing"
)

func TestFindLastIndex(t *testing.T) {
	arr := []TestModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	i := FindLastIndex(arr)
	if i != 2 {
		t.Error("wrong")
	}
}

func TestFindLastIndex_EmptyArray(t *testing.T) {
	if FindLastIndex([]int{}) != -1 {
		t.Error("err")
	}
}

func TestFindLastIndex_NotArray(t *testing.T) {
	if FindLastIndex(nil) != -1 {
		t.Fatal("err")
	}

	dic := map[string]int{
		"a": 1,
		"b": 2,
	}
	if FindLastIndex(dic) != -1 {
		t.Error("err")
	}
}

func TestChain_FindLastIndex(t *testing.T) {
	arr := []TestModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	var index int
	Chain(arr).FindLastIndex().Value(&index)
	if index != 2 {
		t.Error("wrong")
	}
}
