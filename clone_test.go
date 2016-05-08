package underscore

import (
	"testing"
)

func TestClone(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	duplicate := Clone(arr)
	ok := All(duplicate, func (n, i int) bool {
		return arr[i] == n
	})
	if !ok {
		t.Error("wrong")
	}
}

func TestChain_Clone(t *testing.T) {
	dict := map[string]int{ "a": 1, "b": 2 }
	res := Chain(dict).Clone().Value()
	duplicate := res.(map[string]int)
	ok := All(duplicate, func (n int, key string) bool {
		return dict[key] == n
	})
	if !ok {
		t.Error("wrong")
	}
}