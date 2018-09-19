package underscore

import (
	"testing"
)

func TestFindLastIndex(t *testing.T) {
	arr := []TestModel{
		TestModel{1, "one"},
		TestModel{1, "two"},
		TestModel{1, "three"},
	}
	i := FindLastIndex(arr)

	if i != 2 {
		t.Error("wrong")
	}
}

func TestChain_FindLastIndex(t *testing.T) {
	arr := []TestModel{
		TestModel{1, "one"},
		TestModel{1, "two"},
		TestModel{1, "three"},
	}
	res := Chain(arr).FindLastIndex().Value()

	if res.(int) != 2 {
		t.Error("wrong")
	}
}
