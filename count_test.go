package underscore

import (
	"testing"
)

func TestCount(t *testing.T) {
	count := Count(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	if count != 3 {
		t.Error("wrong")
	}

	count = Count([]string{ "a", "b", "c" })
	if count != 3 {
		t.Error("wrong")
	}
}