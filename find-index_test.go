package underscore

import (
	"testing"
)

func TestFindIndex(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 1, Name: "two"},
		TestModel{ID: 1, Name: "three"},
	}
	i := FindIndex(arr, func(r TestModel, _ int) bool {
		return r.Name == arr[1].Name
	})
	if i != 1 {
		t.Error("wrong")
	}
}

func TestFindIndexBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	i := FindIndexBy(arr, map[string]interface{}{
		"id": 1,
	})
	if i != 0 {
		t.Error("wrong")
	}
}

func TestChain_FindIndex(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	index := Chain(arr).FindIndex(func(r TestModel, _ int) bool {
		return r.Name == arr[1].Name
	})
	if index != 1 {
		t.Error("wrong")
	}
}

func TestChain_FindIndexBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	index := Chain(arr).FindIndexBy(map[string]interface{}{
		"id": 1,
	})
	if index == -1 || arr[index].ID != 1 {
		t.Error("wrong")
	}
}
