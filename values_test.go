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
	v, _ := Values(dict)
	res, ok := v.([]string)
	if !(ok && len(res) == len(dict)) {
		t.Error("wrong")
	}
}

func TestChain_Values(t *testing.T) {
	arr := []string{ "a", "b" }
	_, err := Chain(arr).Values().Value()
	if err == nil {
		t.Error("wrong")
	}
}