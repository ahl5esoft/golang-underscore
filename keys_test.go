package underscore

import (
	"testing"
)

func TestKeys(t *testing.T) {
	arr := []string{ "aa" }
	v := Keys(arr)
	if v != nil {
		t.Error("wrong")
		return
	}

	dict := map[int]string{	
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}
	v = Keys(dict)
	res, ok := v.([]int)
	if !(ok && len(res) == len(dict)) {
		t.Error("wrong")
	}
}

func TestChain_Keys(t *testing.T) {
	dict := map[int]string{	
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}
	v := Chain(dict).Keys().Value()
	res, ok := v.([]int)
	if !(ok && len(res) == len(dict)) {
		t.Error("wrong")
	}
}