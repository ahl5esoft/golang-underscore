package underscore

import "testing"

type testModel struct {
	testNestedModel

	ID   int
	Name string
}

type testNestedModel struct {
	Age int
}

func Test_All(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}
	ok := All(arr, func(r testModel, _ int) bool {
		return r.ID == 1
	})
	if !ok {
		t.Error("wrong")
	}
}

func Test_AllBy(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}
	ok := AllBy(arr, nil)
	if ok {
		t.Fatal("wrong")
	}

	ok = AllBy(arr, map[string]interface{}{
		"name": "a",
	})
	if ok {
		t.Fatal("wrong")
	}

	ok = AllBy(arr, map[string]interface{}{
		"id": 1,
	})
	if !ok {
		t.Error("wrong")
	}
}

func Test_Chain_All_False(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}).All(func(r testModel, _ int) bool {
		return r.ID != 1
	})
	if ok {
		t.Error("wrong")
	}
}

func Test_Chain_All_True(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}).All(func(r testModel, _ int) bool {
		return r.ID == 1
	})
	if !ok {
		t.Error("wrong")
	}
}

func Test_Chain_AllBy_False(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).AllBy(map[string]interface{}{
		"Name": "a",
	})
	if ok {
		t.Error("wrong")
	}
}

func Test_Chain_AllBy_True(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "one"},
		{ID: 3, Name: "one"},
	}).AllBy(map[string]interface{}{
		"Name": "one",
	})
	if !ok {
		t.Error("wrong")
	}
}
