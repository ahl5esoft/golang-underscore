package underscore

import (
	"testing"
)

func Test_Reduce(t *testing.T) {
	v := Reduce([]int{1, 2}, func(memo []int, n, _ int) []int {
		memo = append(memo, n)
		memo = append(memo, n+10)
		return memo
	}, make([]int, 0))
	res, ok := v.([]int)
	if !(ok && len(res) == 4) {
		t.Error("wrong length")
	}

	if !(res[0] == 1 && res[1] == 11 && res[2] == 2 && res[3] == 12) {
		t.Error("wrong result")
	}
}

func Test_Chain_Reduce(t *testing.T) {
	res := make([]int, 0)
	Chain([]int{1, 2}).Reduce(func(memo []int, n, _ int) []int {
		memo = append(memo, n)
		memo = append(memo, n+10)
		return memo
	}, make([]int, 0)).Value(&res)
	if !(len(res) == 4 && res[0] == 1 && res[1] == 11 && res[2] == 2 && res[3] == 12) {
		t.Error(res)
	}
}
