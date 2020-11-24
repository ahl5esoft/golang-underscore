package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Skip(t *testing.T) {
	src := []int{1, 2, 3}
	var res []int
	Chain(src).Skip(2).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]int{3},
	)
}
