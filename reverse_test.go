package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reverse(t *testing.T) {
	src := []testModel{
		{ID: 2, Name: "two"},
		{ID: 1, Name: "one"},
		{ID: 3, Name: "three"},
	}
	var res []testModel
	Chain(src).Reverse(func(r testModel, _ int) int {
		return r.ID
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]testModel{
			src[2],
			src[0],
			src[1],
		},
	)
}

func Test_ReverseBy(t *testing.T) {
	src := []testModel{
		{ID: 2, Name: "two"},
		{ID: 1, Name: "one"},
		{ID: 3, Name: "three"},
	}
	var res []testModel
	Chain(src).ReverseBy("ID").Value(&res)
	assert.EqualValues(
		t,
		res,
		[]testModel{
			src[2],
			src[0],
			src[1],
		},
	)
}
