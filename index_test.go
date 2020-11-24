package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Index(t *testing.T) {
	src := []string{"a", "b"}
	res := make(map[string]string)
	Chain(src).Index(func(item string, _ int) string {
		return item
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		map[string]string{
			"a": "a",
			"b": "b",
		},
	)
}

func Test_IndexBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "a"},
		{ID: 2, Name: "a"},
		{ID: 3, Name: "b"},
		{ID: 4, Name: "b"},
	}
	res := make(map[string]testModel)
	Chain(src).IndexBy("name").Value(&res)
	assert.EqualValues(
		t,
		res,
		map[string]testModel{
			"a": src[1],
			"b": src[3],
		},
	)
}
