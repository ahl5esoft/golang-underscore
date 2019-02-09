package underscore

import (
	"testing"
)

func Test_Values_Array(t *testing.T) {
	arr := []int{1, 2, 3}
	res := make([]string, 0)
	Values(arr, &res)
	if len(res) != 0 {
		t.Error(res)
	}
}

func Test_Values_Hash(t *testing.T) {
	dict := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}
	res := make([]string, 0)
	Values(dict, &res)
	if len(res) != len(dict) {
		t.Error(res)
	}
}

func Test_Chain_Values_Array(t *testing.T) {
	arr := []string{"a", "b"}
	res := make([]string, 0)
	Chain(arr).Values().Value(&res)
	if len(res) != len(arr) {
		t.Error(res)
	}
}

func Test_Chain_Values_Map(t *testing.T) {
	dict := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}
	res := make([]string, 0)
	Chain(dict).Values().Value(&res)
	if len(res) != len(dict) {
		t.Error(res)
	}
}
