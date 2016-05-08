package underscore

import (
	"testing"
)

func TestTake(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	v := Take(arr, 1)
	res, ok := v.([]int)
	if !ok {
		t.Error("wrong")
		return
	}

	if res[0] != 1 {
		t.Error("wrong")
	}
}

func TestChain_Take(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	v := Chain(arr).Take(1).Value()
	res, ok := v.([]int)
	if !ok {
		t.Error("wrong")
		return
	}

	if res[0] != 1 {
		t.Error("wrong")
		return
	}

	v = Chain(nil).Take(1).Value()
	if v != nil {
		t.Error("wrong")
	}
}