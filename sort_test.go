package underscore

import (
	"testing"
)

func TestSort(t *testing.T) {
	res, _ := Sort([]int{ 5, 3, 2, 1 }, func (thisValue, _, thatValue, _ interface{}) bool {
		return thisValue.(int) < thatValue.(int)
	})
	if res[0].(int) != 1 {
		t.Error("Sort error")
	}
	
	t.Log(res)
}