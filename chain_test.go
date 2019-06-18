package underscore

import "testing"

var (
	benchmarkSize = 1000000
)

type testModel struct {
	ID   int
	Name string
}

type testNestedModel struct {
	testModel

	Age int
}

func Benchmark_Chain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var dst int
		Range(1, benchmarkSize, 1).Map(func(r, _ int) int {
			return -r
		}).Where(func(r, _ int) bool {
			return r < -20
		}).First().Value(&dst)
	}
}

func Benchmark_Chain_New(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var dst int
		Range2(1, benchmarkSize, 1).Select(func(r, _ int) int {
			return -r
		}).Where(func(r, _ int) bool {
			return r < -20
		}).First().Value(&dst)
	}
}
