package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Object_interface(t *testing.T) {
	src := [][]interface{}{
		{"a", 1},
		{"b", 2},
	}
	res := make(map[string]int)
	Chain(src).Object().Value(&res)
	assert.EqualValues(
		t,
		res,
		map[string]int{
			"a": 1,
			"b": 2,
		},
	)
}

func Test_Object_string(t *testing.T) {
	src := [][]string{
		{"a", "a1"},
		{"b", "b1"},
	}
	res := make(map[string]string)
	Chain(src).Object().Value(&res)
	assert.EqualValues(
		t,
		res,
		map[string]string{
			"a": src[0][1],
			"b": src[1][1],
		},
	)
}
