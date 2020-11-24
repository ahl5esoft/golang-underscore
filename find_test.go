package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_Find(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var dst int
		Range(1, benchmarkSize, 1).Find(func(r, _ int) bool {
			return r > 0
		}).Value(&dst)
	}
}

func Test_Find(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	var res testModel
	Chain(arr).Find(func(r testModel, _ int) bool {
		return r.ID == 1
	}).Value(&res)
	assert.Equal(t, res, arr[0])
}

func Test_FindBy(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	var res testModel
	Chain(arr).FindBy(map[string]interface{}{
		"id": 2,
	}).Value(&res)
	assert.Equal(t, res, arr[1])
}
