package underscore

import "testing"

func Test_Skip(t *testing.T) {
	src := []int{1, 2, 3}
	dst := make([]int, 0)
	Chain2(src).Skip(2).Value(&dst)
	if len(dst) != 1 || dst[0] != 3 {
		t.Fatal(dst)
	}
}
