package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_Any(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Range(1, benchmarkSize, 1).Any(func(r, _ int) bool {
			return r > 1000
		})
	}
}

func Test_Any_False(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).Any(func(r testModel, _ int) bool {
		return r.ID == 0
	})
	assert.False(t, ok)
}

func Test_Any_True(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).Any(func(r testModel, _ int) bool {
		return r.ID == 1
	})
	assert.True(t, ok)
}

func Test_AnyBy_False(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).AnyBy(map[string]interface{}{
		"id": 0,
	})
	assert.False(t, ok)
}

func Test_AnyBy_True(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).AnyBy(map[string]interface{}{
		"name": "two",
	})
	assert.True(t, ok)
}
