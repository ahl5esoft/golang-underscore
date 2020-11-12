package underscore

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_Each(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}
	Chain(arr).Each(func(r testModel, i int) {
		assert.Equal(t, r, arr[i])
	})
}
