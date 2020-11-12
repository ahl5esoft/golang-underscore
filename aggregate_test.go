package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_Aggregate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		total := 0
		Range(1, 100, 1).Aggregate(make([]int, 0), func(memo []int, r, _ int) []int {
			memo = append(memo, r)
			memo = append(memo, -r)
			return memo
		}).Value(&total)
	}
}

func Benchmark_Aggregate_NoValue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Range(1, 100, 1).Aggregate(make([]int, 0), func(memo []int, r, _ int) []int {
			memo = append(memo, r)
			memo = append(memo, -r)
			return memo
		})
	}
}

func Test_Aggregate(t *testing.T) {
	dst := make([]int, 0)
	Chain([]int{1, 2}).Aggregate(make([]int, 0), func(memo []int, n, _ int) []int {
		memo = append(memo, n)
		memo = append(memo, n+10)
		return memo
	}).Value(&dst)
	assert.EqualValues(
		t,
		dst,
		[]int{1, 11, 2, 12},
	)
}
