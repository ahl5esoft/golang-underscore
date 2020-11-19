package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Except(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	var res []int
	Chain(arr).Except(func(n, i int) bool {
		return n%2 == 0
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]int{1, 3},
	)
}

func Test_ExceptBy(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	var res []testModel
	Chain(arr).ExceptBy(map[string]interface{}{
		"Id": 1,
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]testModel{
			{ID: 2, Name: "two"},
			{ID: 3, Name: "three"},
		},
	)
}
