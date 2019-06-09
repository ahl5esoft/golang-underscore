package underscore

import (
	"strconv"
	"testing"
)

func Test_Select(t *testing.T) {
	src := []string{"11", "12", "13"}
	dst := make([]int, 0)
	Chain2(src).Select(func(s string, _ int) int {
		n, _ := strconv.Atoi(s)
		return n
	}).Value(&dst)
	if len(dst) != len(src) {
		t.Error(dst)
	}
}

func Test_SelectBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	dst := make([]string, 0)
	Chain2(src).SelectBy("name").Value(&dst)
	if len(dst) != len(src) {
		t.Fatal(dst)
	}

	for i := 0; i < 3; i++ {
		if dst[i] != src[i].Name {
			t.Error("wrong result")
		}
	}
}
