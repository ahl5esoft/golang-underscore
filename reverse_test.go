package underscore

import (
	"testing"
)

func Test_Reverse(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 3, Name: "three"},
	}
	res := make([]TestModel, 0)
	Reverse(arr, func(n TestModel, _ int) int {
		return n.ID
	}, &res)
	if len(res) != len(arr) {
		t.Fatal("wrong length")
	}

	if !(res[0].ID == 3 && res[1].ID == 2 && res[2].ID == 1) {
		t.Error("reverse error")
	}
}

func Test_ReverseBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 3, Name: "three"},
	}
	res := make([]TestModel, 0)
	ReverseBy(arr, "id", &res)
	if len(res) != len(arr) {
		t.Fatal("wrong length")
	}

	if !(res[0].ID == 3 && res[1].ID == 2 && res[2].ID == 1) {
		t.Error("reverse error")
	}
}

func Test_Chain_Reverse(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 3, Name: "three"},
	}
	res := make([]TestModel, 0)
	Chain(arr).Reverse(func(n TestModel, _ int) int {
		return n.ID
	}).Value(&res)
	if len(res) != len(arr) {
		t.Fatal("wrong length")
	}

	if !(res[0].ID == 3 && res[1].ID == 2 && res[2].ID == 1) {
		t.Error("reverse error")
	}
}

func Test_Chain_ReverseBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 3, Name: "three"},
	}
	res := make([]TestModel, 0)
	Chain(arr).ReverseBy("id").Value(&res)
	if len(res) != len(arr) {
		t.Fatal("wrong length")
	}

	if !(res[0].ID == 3 && res[1].ID == 2 && res[2].ID == 1) {
		t.Error("reverseBy error")
	}
}
