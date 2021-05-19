package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_Aggregate(b *testing.B) {
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			total := 0
			Range(1, 100, 1).Aggregate(
				func(memo []int, r, _ int) []int {
					memo = append(memo, r)
					memo = append(memo, -r)
					return memo
				},
				make([]int, 0),
			).Value(&total)
		}
	})

	b.Run("no value", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Range(1, 100, 1).Aggregate(
				func(memo []int, r, _ int) []int {
					memo = append(memo, r)
					memo = append(memo, -r)
					return memo
				},
				make([]int, 0),
			)
		}
	})
}

func Test_Aggregate(t *testing.T) {
	var res []int
	Chain([]int{1, 2}).Aggregate(
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
