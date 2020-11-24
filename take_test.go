package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_Take(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dst := make([]int, 0)
		Range(1, benchmarkSize, 1).Take(200).Value(&dst)
	}
}

func Benchmark_Take_New(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dst := make([]int, 0)
		Range(1, benchmarkSize, 1).Take(200).Value(&dst)
	}
}

func Test_Take(t *testing.T) {
	src := []int{1, 2, 3}
	var res []int
	Chain(src).Take(1).Value(&res)
	assert.EqualValues(
		t,
		res,
		src[:1],
	)
}
