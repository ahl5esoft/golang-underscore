package underscore

import (
	"testing"
)

func Test_Group(t *testing.T) {
	dic := Group([]int{1, 2, 3, 4, 5}, func(n, _ int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	}).(map[string][]int)
	if len(dic["even"]) != 2 {
		t.Error("wrong")
	}
}

func Test_GroupBy(t *testing.T) {
	arr := []TestModel{
		{ID: 1, Name: "a"},
		{ID: 2, Name: "a"},
		{ID: 3, Name: "b"},
		{ID: 4, Name: "b"},
	}
	dic := GroupBy(arr, "name").(map[string][]TestModel)
	if len(dic) != 2 {
		t.Error("wrong")
	}
}

func TestChain_Group(t *testing.T) {
	res := make(map[string][]int)
	Chain([]int{1, 2, 3, 4, 5}).Group(func(n, _ int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	}).Value(&res)
	if len(res["even"]) != 2 {
		t.Error("wrong")
	}
}

func TestChain_GroupBy(t *testing.T) {
	res := make(map[string][]TestModel)
	Chain([]TestModel{
		{ID: 1, Name: "a"},
		{ID: 2, Name: "a"},
		{ID: 3, Name: "b"},
		{ID: 4, Name: "b"},
	}).GroupBy("Name").Value(&res)
	if len(res) != 2 {
		t.Error("wrong")
	}
}
