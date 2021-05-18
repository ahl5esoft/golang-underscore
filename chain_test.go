package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var benchmarkSize = 1000000

type testModel struct {
	ID   int
	Name string
}

type testNestedModel struct {
	testModel

	Age int
}

type testSelectManyModel struct {
	Str   string
	Slice []string
	Array [2]int
}

func Benchmark_Chain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var dst int
		Range(1, benchmarkSize, 1).Select(func(r, _ int) int {
			return -r
		}).Where(func(r, _ int) bool {
			return r < -20
		}).First().Value(&dst)
	}
}

func Test_Chain(t *testing.T) {
	t.Run("指针", func(t *testing.T) {
		src := []int{1, 2, 3}
		var res []int
		Chain(&src).MapMany(func(r, i int) []int {
			return []int{r, i}
		}).Value(&res)
		assert.EqualValues(t, res, []int{1, 0, 2, 1, 3, 2})
	})
}
