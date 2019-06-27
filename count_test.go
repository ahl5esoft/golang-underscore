package underscore

import "testing"

func Test_Count(t *testing.T) {
	src := []string{"a", "b", "c"}
	dst := Chain(src).Count()
	if dst != len(src) {
		t.Error("wrong")
	}
}
