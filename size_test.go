package underscore

import (
	"testing"
)

func Test_Size(t *testing.T) {
	dict := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	if Size(dict) != len(dict) {
		t.Error("wrong")
	}
}

func Test_Chain_Size(t *testing.T) {
	arr := []string{"a", "b", "c"}
	size := Chain(arr).Size()
	if size != len(arr) {
		t.Error("wrong")
	}
}
