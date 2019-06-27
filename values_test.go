package underscore

import "testing"

func Test_Values_Array(t *testing.T) {
	src := []string{"a", "b"}
	dst := make([]string, 0)
	Chain(src).Values().Value(&dst)
	if len(src) == 0 {
		t.Error(src)
	}
}

func Test_Values_Map(t *testing.T) {
	src := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}
	dst := make([]string, 0)
	Chain(src).Values().Value(&dst)
	if len(src) != len(dst) {
		t.Error(src)
	}
}
