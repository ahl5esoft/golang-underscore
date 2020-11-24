package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Range(t *testing.T) {
	q := Range(0, 100, 1)

	odd := q.Where(func(r, _ int) bool {
		return r%2 == 1
	}).Count()
	assert.Equal(t, odd, 50)

	even := q.Where(func(r, _ int) bool {
		return r%2 == 0
	}).Count()
	assert.Equal(t, even, 50)
}

func Test_Range_StepEq0(t *testing.T) {
	defer func() {
		assert.NotNil(
			t,
			recover(),
		)
	}()

	dst := make([]int, 0)
	Range(0, 10, 0).Value(&dst)
}

func Test_Range_StartEqStop(t *testing.T) {
	var res []int
	Range(0, 0, 1).Value(&res)
	assert.Len(t, res, 0)
}

func Test_Range_Increment(t *testing.T) {
	size := 10
	var res []int
	Range(0, size, 1).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	)
}

func Test_Range_Decrement(t *testing.T) {
	start := 10
	step := -2
	var res []int
	Range(start, 0, step).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]int{10, 8, 6, 4, 2},
	)
}
