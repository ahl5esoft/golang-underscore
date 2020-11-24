package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Keys_Array(t *testing.T) {
	src := []string{"aa", "bb", "cc"}
	var res []int
	Chain(src).Keys().Value(&res)
	assert.EqualValues(
		t,
		res,
		[]int{0, 1, 2},
	)
}

func Test_Keys_Map(t *testing.T) {
	src := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}
	var res []int
	Chain(src).Keys().Sort(func(r, _ int) int {
		return r
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]int{1, 2, 3, 4},
	)
}
