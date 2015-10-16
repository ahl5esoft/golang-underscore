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
	v, err := Chain(arr).Size().Value()
	if err != nil {
		t.Error(err)
	}

	res, ok := v.(int)
	if !(ok && res == len(arr)) {
		t.Error("wrong")
	}
}