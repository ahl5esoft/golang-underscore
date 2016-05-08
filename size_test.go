package underscore

import (
	"testing"
)

func TestSize(t *testing.T) {
	dict := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	if Size(dict) != len(dict) {
		t.Error("wrong")
	}
}

func TestChain_Size(t *testing.T) {
	arr := []string{ "a", "b", "c" }
	v := Chain(arr).Size().Value()
	res, ok := v.(int)
	if !(ok && res == len(arr)) {
		t.Error("wrong")
	}
}