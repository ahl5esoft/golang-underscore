package underscore

import "testing"

func Test_Size(t *testing.T) {
	src := []string{"a", "b", "c"}
	size := Chain(src).Size()
	if size != len(src) {
		t.Error("wrong")
	}
}
