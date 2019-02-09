package underscore

import (
	"strconv"
	"testing"
)

func Test_Map(t *testing.T) {
	arr := []string{"11", "12", "13"}
	res := make([]int, 0)
	Map(arr, func(s string, _ int) int {
		n, _ := strconv.Atoi(s)
		return n
	}, &res)
	if len(res) != len(arr) {
		t.Error(res)
	}
}

func Test_MapBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	res := make([]string, 0)
	MapBy(arr, "name", &res)
	if len(res) != len(arr) {
		t.Fatal(res)
	}

	for i := 0; i < 3; i++ {
		if res[i] != arr[i].Name {
			t.Error("wrong result")
		}
	}
}

func Test_Chain_Map(t *testing.T) {
	arr := []string{"a", "b", "c"}
	res := make([]string, 0)
	Chain(arr).Map(func(item, _ interface{}) string {
		return item.(string) + "-"
	}).Value(&res)
	if !(len(res) == len(arr) && res[0] == "a-") {
		t.Error("wrong")
	}
}

func Test_Chain_MapBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	res := make([]string, 0)
	Chain(arr).MapBy("Name").Value(&res)
	if len(res) != len(arr) {
		t.Fatal("wrong length")
	}

	for i := 0; i < 3; i++ {
		if res[i] != arr[i].Name {
			t.Error("wrong result")
		}
	}
}
