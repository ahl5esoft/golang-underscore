package underscore

import (
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{ 1, 2, 3, 5 }
	res, _ := Sort([]int{ 5, 3, 2, 1 }, func (thisValue, _, thatValue, _ interface{}) bool {
		return thisValue.(int) < thatValue.(int)
	})

	for i, n := range arr {
		if res[i].(int) != n {
			t.Error("wrong result")
		}
	}
}

func TestChain_Sort(t *testing.T) {
	arr := []int{ 1, 2, 3, 5 }
	v, _ := Chain([]int{ 5, 3, 2, 1 }).Sort(func (thisValue, _, thatValue, _ interface{}) bool {
		return thisValue.(int) < thatValue.(int)
	}).Value()

	res, ok := v.([]interface{})
	if !(ok && len(res) == len(arr)) {
		t.Error("wrong length")
	}

	for i, n := range arr {
		if res[i].(int) != n {
			t.Error("wrong result")
		}
	}
}