package underscore

import (
	"testing"
)

func TestLast(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	n, ok := Last(arr).(int)
	if !(ok && n == 3) {
		t.Error("wrong")
		return
	}

	dict := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	str, ok := Last(dict).(string)
	if !(ok && str == "bb") {
		t.Error("wrong")
	}
}

func TestChain_Last(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	n, ok := Chain(arr).Last().Value().(int)
	if !(ok && n == 3) {
		t.Error("wrong")
		return
	}

	dict := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	str, ok := Chain(dict).Last().Value().(string)
	if !(ok && str == "bb") {
		t.Error("wrong")
	}
}