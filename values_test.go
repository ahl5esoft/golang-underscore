package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Values_Array(t *testing.T) {
	src := []string{"a", "b"}
	var res []string
	Chain(src).Values().Value(&res)
	assert.EqualValues(t, res, src)
}

func Test_Values_Map(t *testing.T) {
	src := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}
	var res []string
	Chain(src).Values().Order(func(r string, _ int) string {
		return r
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]string{"a", "b", "c", "d"},
	)
}
