package underscore

import (
	"testing"
)

func TestUniq(t *testing.T) {
	res, _ := Uniq([]int{ 1, 2, 1, 4, 1, 3 })
	if len(res) != 4 {
		t.Error("wrong")
	}
}

func TestChain_Uniq(t *testing.T) {
	v, _ := Chain([]int{ 1, 2, 1, 4, 1, 3 }).Uniq().Value()
	res, ok := v.([]interface{})
	if !(ok && len(res) == 4) {
		t.Error("wrong")
	}
}

func TestUniqBy(t *testing.T) {
	res, _ := UniqBy([]int{ 1, 2, 1, 4, 1, 3 }, func (item interface{}, _ int) interface{} {
		return item.(int) % 2
	})
	if len(res) != 2 {
		t.Error("wrong")
	}
}

func TestChain_UniqBy(t *testing.T) {
	v, _ := Chain([]int{ 1, 2, 1, 4, 1, 3 }).UniqBy(func (item interface{}, _ int) interface{} {
		return item.(int) % 2
	}).Value()
	res, ok := v.([]interface{})
	if !(ok && len(res) == 2) {
		t.Error("wrong")
	}
}