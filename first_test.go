package underscore

import (
	"testing"
)

func TestFirst(t *testing.T) {
	arr := []int{1, 2, 3}
	var n int
	First(arr, &n)
	if n != 1 {
		t.Fatal("wrong")
	}

	First(nil, &n)
	if n != 1 {
		t.Error("wrong")
	}
}

func TestChain_First(t *testing.T) {
	arr := []int{1, 2, 3}
	var item int
	Chain(arr).First().Value(&item)
	if item != 1 {
		t.Error("wrong")
	}
}
