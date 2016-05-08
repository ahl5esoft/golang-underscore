package underscore

import (
	"testing"
)

func TestFirst(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	v := First(arr)
	n, ok := v.(int)
	if !(ok && n == 1) {
		t.Error("wrong")
		return
	}

	v = First(nil)
	if v != nil {
		t.Error("wrong")
	}
}

func TestChain_Test(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	v := Chain(arr).First().Value()
	n, ok := v.(int)
	if !(ok && n == 1) {
		t.Error("wrong")
		return
	}

	v = Chain(nil).First().Value()
	if v != nil {
		t.Error("wrong")
	}
}