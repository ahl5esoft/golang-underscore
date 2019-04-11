package underscore

import (
	"testing"
)

func TestAny(t *testing.T) {
	arr := []TestModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	ok := Any(arr, func(r TestModel, _ int) bool {
		return r.ID == 0
	})
	if ok {
		t.Error("wrong")
	}
}

func TestAnyBy(t *testing.T) {
	arr := []TestModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
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

func TestChain_Any_False(t *testing.T) {
	ok := Chain([]TestModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).Any(func(r TestModel, _ int) bool {
		return r.ID == 0
	})
	if ok {
		t.Error("wrong")
	}
}

func TestChain_Any_True(t *testing.T) {
	ok := Chain([]TestModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).Any(func(r TestModel, _ int) bool {
		return r.ID == 1
	})
	if !ok {
		t.Error("wrong")
	}
}

func TestChain_AnyBy_False(t *testing.T) {
	ok := Chain([]TestModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).AnyBy(map[string]interface{}{
		"id": 0,
	})
	if ok {
		t.Error("wrong")
	}
}

func TestChain_AnyBy_True(t *testing.T) {
	ok := Chain([]TestModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).AnyBy(map[string]interface{}{
		"name": "two",
	})
	if !ok {
		t.Error("wrong")
	}
}
