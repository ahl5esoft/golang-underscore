package underscore

import (
	"testing"
)

func TestAny(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	ok := Any(arr, func(r TestModel, _ int) bool {
		return r.ID == 0
	})
	if ok {
		t.Error("wrong")
	}
}

func TestChain_Any(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	res := Chain(arr).Any(func(r TestModel, _ int) bool {
		return r.ID == 0
	}).Value()
	if res.(bool) {
		t.Error("wrong")
	}
}

func TestAnyBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	ok := AnyBy(arr, map[string]interface{}{
		"Id": 0,
	})
	if ok {
		t.Error("wrong")
		return
	}

	ok = AnyBy(arr, map[string]interface{}{
		"id":   arr[0].ID,
		"name": arr[0].Name,
	})
	if !ok {
		t.Error("wrong")
	}
}

func TestChain_AnyBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	res := Chain(arr).AnyBy(map[string]interface{}{
		"id": 0,
	}).Value()
	if res.(bool) {
		t.Error("wrong")
	}
}
