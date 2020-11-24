package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_First(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var dst int
		Range(1, benchmarkSize, 1).First().Value(&dst)
	}
}

func Test_First(t *testing.T) {
	var res int
	Chain([]int{1, 2, 3}).First().Value(&res)
	assert.Equal(t, res, 1)
}

func Test_First_Twice(t *testing.T) {
	var res int
	Chain([][]int{
		{1, 3, 5, 7},
		{2, 4, 6, 8},
	}).First().First().Value(&res)
	assert.EqualValues(t, res, 1)
}
