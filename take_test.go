package underscore

import (
	"testing"
)

func TestTake(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	v, _ := Take(arr, 1)
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
	v, _ := Chain(arr).Take(1).Value()
	res, ok := v.([]int)
	if !ok {
		t.Error("wrong")
		return
	}

	if res[0] != 1 {
		t.Error("wrong")
		return
	}

	v, _ = Chain(nil).Take(1).Value()
	if v != nil {
		t.Error("wrong")
	}
}

func TestFirst(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	v, _ := First(arr)
	n, ok := v.(int)
	if !(ok && n == 1) {
		t.Error("wrong")
		return
	}

	v, _ = First(nil)
	if v != nil {
		t.Error("wrong")
	}
}

func TestChain_Test(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	v, _ := Chain(arr).First().Value()
	n, ok := v.(int)
	if !(ok && n == 1) {
		t.Error("wrong")
		return
	}

	v, _ = Chain(nil).First().Value()
	if v != nil {
		t.Error("wrong")
	}
}