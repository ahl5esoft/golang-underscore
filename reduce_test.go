package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reduce(t *testing.T) {
	var res []int
	Chain([]int{1, 2}).Reduce(
		func(memo []int, n, _ int) []int {
			memo = append(memo, n)
			memo = append(memo, n+10)
			return memo
		},
		make([]int, 0),
	).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]int{1, 11, 2, 12},
	)
}
