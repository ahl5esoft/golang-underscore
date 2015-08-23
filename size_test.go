package underscore

import (
	"testing"
)

func TestSize(t *testing.T) {
	length := Size(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	if length != 3 {
		t.Error("wrong")
	}

	length = Size([]string{ "a", "b", "c" })
	if length != 3 {
		t.Error("wrong")
	}
}