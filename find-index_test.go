package underscore

import (
	"testing"
)

func TestFindIndex(t *testing.T) {
	arr := []TestModel{
		TestModel{1, "one"},
		TestModel{1, "two"},
		TestModel{1, "three"},
	}
	i := FindIndex(arr, func(r TestModel, _ int) bool {
		return r.Name == arr[1].Name
	})
	if i != 1 {
		t.Error("wrong")
	}
}

func TestChain_FindIndex(t *testing.T) {
	arr := []TestModel{
		TestModel{1, "one"},
		TestModel{1, "two"},
		TestModel{1, "three"},
	}
	res := Chain(arr).FindIndex(func(r TestModel, _ int) bool {
		return r.Name == arr[1].Name
	}).Value()
	if res.(int) != 1 {
		t.Error("wrong")
	}
}

func TestFindIndexBy(t *testing.T) {
	arr := []TestModel{
		TestModel{1, "one"},
		TestModel{2, "two"},
		TestModel{3, "three"},
	}
	i := FindIndexBy(arr, map[string]interface{}{
		"id": 1,
	})
	if i != 0 {
		t.Error("wrong")
	}
}

func TestChain_FindIndexBy(t *testing.T) {
	arr := []TestModel{
		TestModel{1, "one"},
		TestModel{2, "two"},
		TestModel{3, "three"},
	}
	res := Chain(arr).FindIndexBy(map[string]interface{}{
		"id": 1,
	}).Value()
	i := res.(int)
	if i == -1 || arr[i].ID != 1 {
		t.Error("wrong")
	}
}
