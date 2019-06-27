package underscore

import "testing"

func Test_Reduce(t *testing.T) {
	dst := make([]int, 0)
	Chain([]int{1, 2}).Reduce(make([]int, 0), func(memo []int, n, _ int) []int {
		memo = append(memo, n)
		memo = append(memo, n+10)
		return memo
	}).Value(&dst)
	if !(len(dst) == 4 && dst[0] == 1 && dst[1] == 11 && dst[2] == 2 && dst[3] == 12) {
		t.Error(dst)
	}
}
