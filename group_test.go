package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_Group(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dst := make([]int, 0)
		Range(1, benchmarkSize, 1).Group(func(n, _ int) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		}).Value(&dst)
	}
}

func Benchmark_Group_New(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dst := make([]int, 0)
		Range(1, benchmarkSize, 1).Group(func(n, _ int) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		}).Value(&dst)
	}
}

func Test_Group(t *testing.T) {
	res := make(map[string][]int)
	Chain([]int{1, 2, 3, 4, 5}).Group(func(n, _ int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		map[string][]int{
			"odd":  {1, 3, 5},
			"even": {2, 4},
		},
	)
}

func Test_GroupBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "a"},
		{ID: 2, Name: "a"},
		{ID: 3, Name: "b"},
		{ID: 4, Name: "b"},
	}
	res := make(map[string][]testModel)
	Chain(src).GroupBy("Name").Value(&res)
	assert.EqualValues(
		t,
		res,
		map[string][]testModel{
			"a": {src[0], src[1]},
			"b": {src[2], src[3]},
		},
	)
}
