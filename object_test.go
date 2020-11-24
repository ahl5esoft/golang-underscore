package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Object(t *testing.T) {
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
