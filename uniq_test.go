package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Uniq(t *testing.T) {
	src := []int{1, 2, 1, 4, 1, 3}
	var res []int
	Chain(src).Uniq(func(n, _ int) (int, error) {
		return n % 2, nil
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]int{1, 2},
	)
}

func Test_Uniq_SelectorIsNil(t *testing.T) {
	src := []int{1, 2, 1, 4, 1, 3}
	var res []int
	Chain(src).Uniq(nil).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]int{1, 2, 4, 3},
	)
}

func Test_UniqBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "a"},
		{ID: 2, Name: "a"},
		{ID: 3, Name: "a"},
	}
	var res []testModel
	Chain(src).UniqBy("name").Value(&res)
	assert.EqualValues(
		t,
		res,
		[]testModel{src[0]},
	)
}
