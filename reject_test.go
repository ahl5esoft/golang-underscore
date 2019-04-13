package underscore

import (
	"testing"
)

func Test_Reject(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	res := Reject(arr, func(n, i int) bool {
		return n%2 == 0
	}).([]int)
	if len(res) != 2 {
		t.Fatal("wrong length")
	}

	if !(res[0] == 1 && res[1] == 3) {
		t.Error("wrong result")
	}
}

func Test_RejectBy(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	res := RejectBy(arr, map[string]interface{}{
		"Id": 1,
	}).([]testModel)
	if len(res) != 2 {
		t.Error("wrong result")
	}
}

func Test_Chain_Reject(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	res := make([]int, 0)
	Chain(arr).Reject(func(n, i int) bool {
		return n%2 == 0
	}).Value(&res)
	if !(len(res) == 2 && res[0] == 1 && res[1] == 3) {
		t.Error("wrong result")
	}
}

func Test_Chain_RejectBy(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	res := make([]testModel, 0)
	Chain(arr).RejectBy(map[string]interface{}{
		"Id": 1,
	}).Value(&res)
	if len(res) != 2 {
		t.Error("wrong result")
	}
}
