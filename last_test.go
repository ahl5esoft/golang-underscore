package underscore

import (
	"testing"
)

func Test_Last_Array(t *testing.T) {
	arr := []int{1, 2, 3}
	var n int
	Last(arr, &n)
	if n != 3 {
		t.Error(n)
	}
}

func Test_Last_Map(t *testing.T) {
	dict := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	var str string
	Last(dict, &str)
	if !(str == "aa" || str == "bb") {
		t.Error(str)
	}
}

func TestChain_Last_Array(t *testing.T) {
	var item int
	Chain([]int{1, 2, 3}).Last().Value(&item)
	if item != 3 {
		t.Error(item)
	}
}

func TestChain_Last_Map(t *testing.T) {
	var item string
	dict := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	Chain(dict).Last().Value(&item)
	if !(item == "aa" || item == "bb") {
		t.Error(item)
	}
}
