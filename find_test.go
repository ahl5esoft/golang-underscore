package underscore

import (
	"testing"
)

func TestFind(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	item := Find(arr, func(r TestModel, _ int) bool {
		return r.ID == 1
	})
	if item != arr[0] {
		t.Error("wrong")
	}
}

func TestChain_Find(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	item := Chain(arr).Find(func(r TestModel, _ int) bool {
		return r.ID == 1
	}).Value()
	if item.(TestModel) != arr[0] {
		t.Error("wrong")
	}
}

func TestFindBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	res := FindBy(arr, map[string]interface{}{
		"id": 2,
	})
	if res.(TestModel) != arr[1] {
		t.Error("wrong")
	}
}

func TestChain_FindBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	res := Chain(arr).FindBy(map[string]interface{}{
		"id": 2,
	}).Value()
	if res.(TestModel) != arr[1] {
		t.Error("wrong")
	}
}
