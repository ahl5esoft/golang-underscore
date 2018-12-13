package underscore

import (
	"testing"
)

func TestChain(t *testing.T) {
	res, ok := Chain([]int{1, 2, 1, 4, 1, 3}).Uniq(nil).Group(func(n, _ int) string {
		if n%2 == 0 {
			return "even"
		}

		return "old"
	}).Value().(map[string][]int)
	if !(ok && len(res) == 2) {
		t.Error("wrong")
	}
}

func TestChain_Chain_DefaultValue(t *testing.T) {
	res, ok := Chain([]int{1, 2, 1, 4, 1, 3}).Uniq(nil).Group(func(n, _ int) string {
		if n%2 == 0 {
			return "even"
		}

		return "old"
	}).ValueOrDefault(make(map[string][]int)).(map[string][]int)
	if !(ok && len(res) == 2) {
		t.Error("wrong")
	}
}

func TestChain_ValueOrDefault_Default(t *testing.T) {
	res, ok := Chain([]int{}).GroupBy("unknow").ValueOrDefault(make(map[string][]int)).(map[string][]int)
	if !(ok && len(res) == 0) {
		t.Error("wrong")
	}
}
