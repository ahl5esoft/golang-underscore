package underscore

import "testing"

func Test_SelectMany(t *testing.T) {
	src := [2]int{1, 2}
	var dst []int
	Chain(src).SelectMany(func(r, _ int) []int {
		return []int{r - 1, r + 1}
	}).Value(&dst)
	if !(len(dst) == 4 && dst[0] == 0 && dst[1] == 2 && dst[2] == 1 && dst[3] == 3) {
		t.Error(dst)
	}
}

func Test_SelectManyBy_Array(t *testing.T) {
	src := []testSelectManyModel{
		{Array: [2]int{1, 2}},
		{Array: [2]int{3, 4}},
	}
	var dst []int
	Chain(src).SelectManyBy("Array").Value(&dst)
	if !(len(dst) == 4 && dst[0] == 1 && dst[1] == 2 && dst[2] == 3 && dst[3] == 4) {
		t.Error(dst)
	}
}

func Test_SelectManyBy_Slice(t *testing.T) {
	src := []testSelectManyModel{
		{Slice: []string{"a", "b"}},
		{Slice: []string{"c", "d"}},
	}
	var dst []string
	Chain(src).SelectManyBy("Slice").Value(&dst)
	if !(len(dst) == 4 && dst[0] == "a" && dst[1] == "b" && dst[2] == "c" && dst[3] == "d") {
		t.Error(dst)
	}
}
