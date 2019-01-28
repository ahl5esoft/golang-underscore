package underscore

import "testing"

func Test_Chain_Value(t *testing.T) {
	res := make(map[string][]int)
	Chain([]int{1, 2, 1, 4, 1, 3}).Uniq(nil).Group(func(n, _ int) string {
		if n%2 == 0 {
			return "even"
		}

		return "old"
	}).Value(&res)
	if len(res) != 2 {
		t.Error("wrong")
	}
}
