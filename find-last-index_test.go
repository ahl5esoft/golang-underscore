package underscore

import (
	"testing"
)

func Test_FindLastIndex(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	i := FindLastIndex(arr)
	if i != 2 {
		t.Error("wrong")
	}
}

func Test_FindLastIndex_EmptyArray(t *testing.T) {
	if FindLastIndex([]int{}) != -1 {
		t.Error("err")
	}
}

func Test_FindLastIndex_NotArray(t *testing.T) {
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

func Test_Chain_FindLastIndex(t *testing.T) {
	arr := []testModel{
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
