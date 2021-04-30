package underscore

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		src := []string{"11", "12", "13"}
		var res []int
		Chain(src).Map(func(s string, _ int) int {
			n, _ := strconv.Atoi(s)
			return n
		}).Value(&res)
		assert.EqualValues(
			t,
			res,
			[]int{11, 12, 13},
		)
	})

	t.Run("ptr", func(t *testing.T) {
		src := []string{"11", "12", "13"}
		var res []*string
		Chain(src).Map(func(r string, _ int) *string {
			return &r
		}).Value(&res)
		assert.EqualValues(
			t,
			res,
			[]*string{&src[0], &src[1], &src[2]},
		)
	})
}

func Test_MapBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	var res []string
	Chain(src).MapBy("name").Value(&res)
	assert.EqualValues(
		t,
		res,
		[]string{"one", "two", "three"},
	)
}
