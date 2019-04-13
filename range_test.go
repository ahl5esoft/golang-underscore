package underscore

import (
	"testing"
)

func Test_Range(t *testing.T) {
	arr := make([]int, 0)
	Range(0, 0, 1).Value(&arr)
	if len(arr) != 0 {
		t.Fatal("wrong")
	}

	Range(0, 10, 0).Value(&arr)
	if len(arr) != 0 {
		t.Fatal("wrong")
	}

	Range(10, 0, 1).Value(&arr)
	if len(arr) != 0 {
		t.Fatal("wrong")
	}

	Range(0, 2, 1).Value(&arr)
	if !(len(arr) == 2 && arr[0] == 0 && arr[1] == 1) {
		t.Fatal("wrong")
	}

	Range(0, 3, 2).Value(&arr)
	if !(len(arr) == 2 && arr[0] == 0 && arr[1] == 2) {
		t.Error("wrong")
	}
}
