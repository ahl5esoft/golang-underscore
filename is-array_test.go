package underscore

import (
	"testing"
)

func Test_IsArray(t *testing.T) {
	if !IsArray([]int{}) {
		t.Error("wrong")
		return
	}

	if IsArray(map[string]int{}) {
		t.Error("wrong")
	}
}
