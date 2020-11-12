package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_All(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Range(0, benchmarkSize, 1).All(func(r, _ int) bool {
			return r < 1000
		})
	}
}

func Test_All_False(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}).All(func(r testModel, _ int) bool {
		return r.Name == "one"
	})
	assert.False(t, ok)
}

func Test_All_True(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}).All(func(r testModel, _ int) bool {
		return r.ID == 1
	})
	assert.True(t, ok)
}

func Test_AllBy_False(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).AllBy(map[string]interface{}{
		"Name": "a",
	})
	assert.False(t, ok)
}

func Test_AllBy_True(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "one"},
		{ID: 3, Name: "one"},
	}).AllBy(map[string]interface{}{
		"name": "one",
	})
	assert.True(t, ok)
}
