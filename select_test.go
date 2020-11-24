package underscore

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Select(t *testing.T) {
	src := []string{"11", "12", "13"}
	var res []int
	Chain(src).Select(func(s string, _ int) int {
		n, _ := strconv.Atoi(s)
		return n
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]int{11, 12, 13},
	)
}

func Test_SelectBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	var res []string
	Chain(src).SelectBy("name").Value(&res)
	assert.EqualValues(
		t,
		res,
		[]string{"one", "two", "three"},
	)
}
