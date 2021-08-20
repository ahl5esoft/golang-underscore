package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Order(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		arr := []testModel{
			{ID: 2, Name: "two"},
			{ID: 1, Name: "one"},
			{ID: 3, Name: "three"},
		}
		var res []int
		Chain(arr).Order(func(n testModel, _ int) int {
			return n.ID
		}).Map(func(r testModel, _ int) int {
			return r.ID
		}).Value(&res)
		assert.Len(
			t,
			res,
			len(arr),
		)
		assert.EqualValues(
			t,
			res,
			[]int{1, 2, 3},
		)
	})

	t.Run("chain", func(t *testing.T) {
		arr := []testModel{
			{ID: 2, Name: "two"},
			{ID: 1, Name: "one"},
			{ID: 3, Name: "three"},
		}
		var res []int
		Chain(arr).Where(func(r testModel, _ int) bool {
			return r.ID > 5
		}).Order(func(r testModel, _ int) int {
			return r.ID
		}).Map(func(r testModel, _ int) int {
			return r.ID
		}).Value(&res)
	})
}

func Test_OrderBy(t *testing.T) {
	arr := []testModel{
		{ID: 2, Name: "two"},
		{ID: 1, Name: "one"},
		{ID: 3, Name: "three"},
	}
	var res []string
	Chain(arr).OrderBy("id").Map(func(r testModel, _ int) string {
		return r.Name
	}).Value(&res)
	assert.Len(t, res, 3)
	assert.EqualValues(
		t,
		res,
		[]string{"one", "two", "three"},
	)
}
