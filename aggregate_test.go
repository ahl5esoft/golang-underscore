package underscore

import (
	"testing"
)

func Benchmark_Aggregate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		total := 0
		Range2(1, 100, 1).Aggregate(make([]int, 0), func(memo []int, r, _ int) []int {
			memo = append(memo, r)
			memo = append(memo, -r)
			return memo
		}).Value(&total)
	}
}

func Benchmark_Aggregate_NoValue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Range2(1, 100, 1).Aggregate(make([]int, 0), func(memo []int, r, _ int) []int {
			memo = append(memo, r)
			memo = append(memo, -r)
			return memo
		})
	}
}

func Test_Aggregate(t *testing.T) {
	dst := make([]int, 0)
	Chain2([]int{1, 2}).Aggregate(make([]int, 0), func(memo []int, n, _ int) []int {
		memo = append(memo, n)
		memo = append(memo, n+10)
		return memo
	}).Value(&dst)
	if !(len(dst) == 4 && dst[0] == 1 && dst[1] == 11 && dst[2] == 2 && dst[3] == 12) {
		t.Error(dst)
	}
}
