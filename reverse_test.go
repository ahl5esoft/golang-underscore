package underscore

import "testing"

func Test_Reverse(t *testing.T) {
	arr := []testModel{
		testModel{ID: 2, Name: "two"},
		testModel{ID: 1, Name: "one"},
		testModel{ID: 3, Name: "three"},
	}
	res := Reverse(arr, func(n testModel, _ int) int {
		return n.ID
	}).([]testModel)
	if len(res) != len(arr) {
		t.Fatal("wrong length")
	}

	if !(res[0].ID == 3 && res[1].ID == 2 && res[2].ID == 1) {
		t.Error("reverse error")
	}
}

func Test_ReverseBy(t *testing.T) {
	arr := []testModel{
		testModel{ID: 2, Name: "two"},
		testModel{ID: 1, Name: "one"},
		testModel{ID: 3, Name: "three"},
	}
	res := ReverseBy(arr, "id").([]testModel)
	if len(res) != len(arr) {
		t.Fatal("wrong length")
	}

	if !(res[0].ID == 3 && res[1].ID == 2 && res[2].ID == 1) {
		t.Error("reverseBy error")
	}
}

func Test_Chain_Reverse(t *testing.T) {
	arr := []testModel{
		testModel{ID: 2, Name: "two"},
		testModel{ID: 1, Name: "one"},
		testModel{ID: 3, Name: "three"},
	}
	var res []testModel
	Chain(arr).Reverse(func(n testModel, _ int) int {
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
	arr := []testModel{
		testModel{ID: 2, Name: "two"},
		testModel{ID: 1, Name: "one"},
		testModel{ID: 3, Name: "three"},
	}
	var res []testModel
	Chain(arr).ReverseBy("id").Value(&res)
	if len(res) != len(arr) {
		t.Fatal("wrong length")
	}

	if !(res[0].ID == 3 && res[1].ID == 2 && res[2].ID == 1) {
		t.Error("reverseBy error")
	}
}
