package underscore

import (
	"testing"
)

func TestGroup(t *testing.T) {
	dict, _ := Group([]int{ 1, 2, 3, 4, 5 }, func (item interface{}, _ interface{}) (interface{}, error) {
		if item.(int) % 2 == 0 {
			return "even", nil
		}
		return "odd", nil
	})
	group, ok := dict["even"]
	if !(ok && len(group) == 2) {
		t.Error("wrong")
	}
}