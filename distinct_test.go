package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Distinct(t *testing.T) {
	src := []int{1, 2, 1, 4, 1, 3}
	dst := make([]int, 0)
	Chain(src).Distinct(func(n, _ int) (int, error) {
		return n % 2, nil
	}).Value(&dst)
	assert.Len(t, dst, 2)
}

func Test_Distinct_SelectorIsNil(t *testing.T) {
	src := []int{1, 2, 1, 4, 1, 3}
	dst := make([]int, 0)
	Chain(src).Distinct(nil).Value(&dst)
	assert.Len(t, dst, 4)
}

func Test_DistinctBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "a"},
		{ID: 2, Name: "a"},
		{ID: 3, Name: "a"},
	}
	dst := make([]testModel, 0)
	Chain(src).DistinctBy("name").Value(&dst)
	assert.Len(t, dst, 1)
}
