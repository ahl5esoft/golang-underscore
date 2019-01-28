package underscore

import (
	"testing"
)

func TestValues(t *testing.T) {
	dict := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}
	v := Values(dict)
	res, ok := v.([]string)
	if !(ok && len(res) == len(dict)) {
		t.Error("wrong")
	}
}

func TestChain_Values_Array(t *testing.T) {
	arr := []string{"a", "b"}
	res := make([]string, 0)
	Chain(arr).Values().Value(&res)
	if res == nil || len(res) != 0 {
		t.Error(res)
	}
}

func TestChain_Values_Map(t *testing.T) {
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
