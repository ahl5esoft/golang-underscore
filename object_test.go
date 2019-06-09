package underscore

import (
	"testing"
)

func Test_Object(t *testing.T) {
	src := [][]interface{}{
		{"a", 1},
		{"b", 2},
	}
	dst := make(map[string]int)
	Chain2(src).Object().Value(&dst)
	if len(dst) != 2 {
		t.Fatal(dst)
	}

	if v, ok := dst["a"]; !(ok && v == 1) {
		t.Fatal("key 1")
	}

	if v, ok := dst["b"]; !(ok && v == 2) {
		t.Error("key 2")
	}
}
