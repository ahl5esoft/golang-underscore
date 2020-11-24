package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SelectMany(t *testing.T) {
	src := [2]int{1, 2}
	var res []int
	Chain(src).SelectMany(func(r, _ int) []int {
		return []int{r - 1, r + 1}
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]int{0, 2, 1, 3},
	)
}

func Test_SelectManyBy_Array(t *testing.T) {
	src := []testSelectManyModel{
		{Array: [2]int{1, 2}},
		{Array: [2]int{3, 4}},
	}
	var res []int
	Chain(src).SelectManyBy("Array").Value(&res)
	assert.EqualValues(
		t,
		res,
		[]int{1, 2, 3, 4},
	)
}

func Test_SelectManyBy_Slice(t *testing.T) {
	src := []testSelectManyModel{
		{Slice: []string{"a", "b"}},
		{Slice: []string{"c", "d"}},
	}
	var res []string
	Chain(src).SelectManyBy("Slice").Value(&res)
	assert.EqualValues(
		t,
		res,
		[]string{"a", "b", "c", "d"},
	)
}
