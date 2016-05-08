package underscore

import (
	"testing"
)

func TestRange(t *testing.T) {
	arr := Range(0, 0, 1)
	if len(arr) != 0 {
		t.Error("wrong")
		return
	}

	arr = Range(0, 10, 0)
	if len(arr) != 0 {
		t.Error("wrong")
		return
	}

	arr = Range(10, 0, 1)
	if len(arr) != 0 {
		t.Error("wrong")
		return
	}

	arr = Range(0, 2, 1)
	if !(len(arr) == 2 && arr[0] == 0 && arr[1] == 1) {
		t.Error("wrong")
		return
	}

	arr = Range(0, 3, 2)
	if !(len(arr) == 2 && arr[0] == 0 && arr[1] == 2) {
		t.Error("wrong")
	}
}

func TestChain_Range(t *testing.T) {
	v := Chain(nil).Range(0, 10, 0).Value()
	arr, ok := v.([]int)
	if !(ok && len(arr) == 0) {
		t.Error("wrong")
		return
	}
	
	v = Chain(nil).Range(0, 3, 1).Value()
	arr, ok = v.([]int)
	if !(ok && len(arr) == 3) {
		t.Error("wrong")
	}
}