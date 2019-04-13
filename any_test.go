package underscore

import (
	"testing"
)

func Test_Any(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	ok := Any(src, func(r testModel, _ int) bool {
		return r.ID == 0
	})
	if ok {
		t.Error("wrong")
	}
}

func Test_AnyBy(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	ok := AnyBy(arr, map[string]interface{}{
		"Id": 0,
	})
	if ok {
		t.Fatal("wrong")
	}

	ok = AnyBy(arr, map[string]interface{}{
		"id":   arr[0].ID,
		"name": arr[0].Name,
	})
	if !ok {
		t.Error("wrong")
	}
}

func Test_Chain_Any_False(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).Any(func(r testModel, _ int) bool {
		return r.ID == 0
	})
	if ok {
		t.Error("wrong")
	}
}

func Test_Chain_Any_True(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).Any(func(r testModel, _ int) bool {
		return r.ID == 1
	})
	if !ok {
		t.Error("wrong")
	}
}

func Test_Chain_AnyBy_False(t *testing.T) {
	ok := Chain([]testModel{
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

func Test_Chain_AnyBy_True(t *testing.T) {
	ok := Chain([]testModel{
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
