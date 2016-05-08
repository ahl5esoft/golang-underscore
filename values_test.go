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

func TestChain_Values(t *testing.T) {
	arr := []string{ "a", "b" }
	res := Chain(arr).Values().Value()
	if res != nil {
		t.Error("wrong")
	}
}