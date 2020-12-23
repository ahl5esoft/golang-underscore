package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_FindIndex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Range(1, benchmarkSize, 1).FindIndex(func(r, _ int) bool {
			return r == 200
		})
	}
}

func Test_FindIndex(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	res := Chain(src).FindIndex(func(r testModel, _ int) bool {
		return r.Name == src[1].Name
	})
	assert.Equal(t, res, 1)
}

func Test_FindIndex_NotExists(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	res := Chain(src).FindIndex(func(r testModel, _ int) bool {
		return r.Name == ""
	})
	assert.Equal(t, res, -1)
}

func Test_FindIndexBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	res := Chain(src).FindIndexBy(map[string]interface{}{
		"id": 1,
	})
	assert.Equal(t, res, 0)
}

func Test_FindIndexBy_NotExists(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	res := Chain(src).FindIndexBy(map[string]interface{}{
		"id": 0,
	})
	assert.Equal(t, res, -1)
}
