package underscore

import (
	"testing"
)

func TestReject(t *testing.T) {
	arr := []int{ 1, 2, 3, 4 }
	v := Reject(arr, func (n, i int) bool {
		return n % 2 == 0
	})
	res, ok := v.([]int)
	if !(ok && len(res) == 2) {
		t.Error("wrong length")
		return
	}

	if !(res[0] == 1 && res[1] == 3) {
		t.Error("wrong result")
	}
}

func TestChain_Reject(t *testing.T) {
	arr := []int{ 1, 2, 3, 4 }
	v := Chain(arr).Reject(func (n, i int) bool {
		return n % 2 == 0
	}).Value()
	res, ok := v.([]int)
	if !(ok && len(res) == 2) {
		t.Error("wrong length")
		return
	}

	if !(res[0] == 1 && res[1] == 3) {
		t.Error("wrong result")
	}
}

func TestRejectBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	v := RejectBy(arr, map[string]interface{}{
		"Id": 1,
	})
	res, ok := v.([]TestModel)
	if !(ok && len(res) == 2) {
		t.Error("wrong result")
	}
}

func TestChain_RejectBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	v := Chain(arr).RejectBy(map[string]interface{}{
		"Id": 1,
	}).Value()
	res, ok := v.([]TestModel)
	if !(ok && len(res) == 2) {
		t.Error("wrong result")
	}
}