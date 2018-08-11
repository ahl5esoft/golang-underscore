package underscore

import (
	"testing"
)

func TestFind(t *testing.T) {
	arr := []TestModel{
		TestModel{1, "one"},
		TestModel{2, "two"},
		TestModel{3, "three"},
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
		TestModel{1, "one"},
		TestModel{2, "two"},
		TestModel{3, "three"},
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
		TestModel{1, "one"},
		TestModel{2, "two"},
		TestModel{3, "three"},
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
		TestModel{1, "one"},
		TestModel{2, "two"},
		TestModel{3, "three"},
	}
	res := Chain(arr).FindBy(map[string]interface{}{
		"id": 2,
	}).Value()
	if res.(TestModel) != arr[1] {
		t.Error("wrong")
	}
}
