package underscore

import (
	"testing"
)

func TestUniq(t *testing.T) {
	res, _ := Uniq([]int{ 1, 2, 1, 4, 1, 3 })
	if len(res) != 4 {
		t.Error("Uniq error")
	}
}

func TestUniqBy(t *testing.T) {
	res, _ := UniqBy([]int{ 1, 2, 1, 4, 1, 3 }, func (item interface{}, _ int) interface{} {
		return item.(int) % 2
	})
	if len(res) != 2 {
		t.Error("UniqBy error")
	}
}