package underscore

import "testing"

func Test_Keys_Array(t *testing.T) {
	src := []string{"aa", "bb", "cc"}
	dst := make([]int, 0)
	Chain2(src).Keys().Value(&dst)
	if len(dst) != len(src) {
		t.Fatal(dst)
	}

	if dst[0] != 0 {
		t.Error(dst)
	}
}

func Test_Keys_Map(t *testing.T) {
	src := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}
	dst := make([]int, 0)
	Chain2(src).Keys().Value(&dst)
	if len(dst) != len(src) {
		t.Fatal(dst)
	}
}
