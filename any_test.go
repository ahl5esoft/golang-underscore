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

func Test_Any(t *testing.T) {
	t.Run("source is empty slice", func(t *testing.T) {
		arr := make([]testModel, 0)
		res := Chain(arr).Any(func(r testModel, _ int) bool {
			return r.ID == 0
		})
		assert.False(t, res)
	})

	t.Run("source is nil", func(t *testing.T) {
		var arr []testModel
		assert.Nil(t, arr)

		res := Chain(arr).Any(func(r testModel, _ int) bool {
			return r.ID == 0
		})
		assert.False(t, res)
	})
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
