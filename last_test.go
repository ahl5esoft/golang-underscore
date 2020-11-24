package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Last(t *testing.T) {
	var res int
	Chain([]int{1, 2, 3}).Last().Value(&res)
	assert.Equal(t, res, 3)
}

func Test_Last_ContinueIterate(t *testing.T) {
	var res []int
	src := [][]int{
		{1, 2, 3, 4},
		{5, 6},
	}
	Chain(src).Last().Map(func(r, _ int) int {
		return r + 5
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]int{10, 11},
	)
}
