package underscore

import (
	"testing"
)

func TestChain(t *testing.T) {
	res, ok := Chain([]int{1, 2, 1, 4, 1, 3}).Uniq(nil).Group(func(n, _ int) string {
		if n%2 == 0 {
			return "even"
		}

		return "old"
	}).Value().(map[string][]int)
	if !(ok && len(res) == 2) {
		t.Error("wrong")
	}
}
