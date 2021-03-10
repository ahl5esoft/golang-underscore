package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Value(t *testing.T) {
	res := make(map[string][]int)
	Chain([]int{1, 2, 1, 4, 1, 3}).Uniq(nil).Group(func(n, _ int) string {
		if n%2 == 0 {
			return "even"
		}

		return "odd"
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		map[string][]int{
			"odd":  {1, 3},
			"even": {2, 4},
		},
	)
}
