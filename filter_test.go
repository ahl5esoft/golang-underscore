package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Filter(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "one"},
		{ID: 3, Name: "three"},
		{ID: 4, Name: "three"},
	}
	var res []testModel
	Chain(src).Filter(func(r testModel, _ int) bool {
		return r.ID%2 == 0
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]testModel{src[1], src[3]},
	)
}

func Test_FilterBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "one"},
		{ID: 3, Name: "three"},
		{ID: 4, Name: "three"},
	}
	var res []testModel
	Chain(src).FilterBy(map[string]interface{}{
		"Name": "one",
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]testModel{src[0], src[1]},
	)
}
