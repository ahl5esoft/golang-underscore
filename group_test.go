package underscore

import (
	"testing"
)

func TestGroup(t *testing.T) {
	dict := Group([]int{ 1, 2, 3, 4, 5 }, func (item interface{}) interface{} {
		if item.(int) % 2 == 0 {
			return "even"
		}
		return "odd"
	})
	group, ok := dict["even"]
	if !(ok && len(group) == 2) {
		t.Error("wrong")
	}
}