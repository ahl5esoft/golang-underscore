package underscore

import (
	"testing"
)

func Test_Take(t *testing.T) {
	arr := []int{1, 2, 3}
	res := Take(arr, 1).([]int)
	if len(res) != 1 || res[0] != 1 {
		t.Error("wrong")
	}
}

func Test_Chain_Take(t *testing.T) {
	arr := []int{1, 2, 3}
	res := make([]int, 0)
	Chain(arr).Take(1).Value(&res)
	if res[0] != 1 {
		t.Fatal("wrong")
	}

	res = make([]int, 0)
	Chain(nil).Take(1).Value(&res)
	if len(res) > 0 {
		t.Error("wrong")
	}
}
