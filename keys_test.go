package underscore

import (
	"testing"
)

func Test_Keys_Array(t *testing.T) {
	arr := []string{"aa"}
	res := make([]string, 0)
	Keys(arr, &res)
	if len(res) > 0 {
		t.Error(res)
	}
}

func Test_Keys_Hash(t *testing.T) {
	dict := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}
	res := make([]int, 0)
	Keys(dict, &res)
	if len(res) != len(dict) {
		t.Error(res)
	}
}

func Test_Chain_Keys_Array(t *testing.T) {
	arr := []string{"aa"}
	res := make([]string, 0)
	Chain(arr).Keys().Value(&res)
	if len(res) != len(arr) {
		t.Error(res)
	}
}

func Test_Chain_Keys_Hash(t *testing.T) {
	dict := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}
	res := make([]int, 0)
	Chain(dict).Keys().Value(&res)
	if len(res) != len(dict) {
		t.Error("wrong")
	}
}
