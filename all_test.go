package underscore

import "testing"

// TestModel is 测试模型
type TestModel struct {
	TestNestedModel

	ID   int
	Name string
}

// TestNestedModel is 嵌套模型
type TestNestedModel struct {
	Age int
}

func TestAll(t *testing.T) {
	arr := []TestModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}
	ok := All(arr, func(r TestModel, _ int) bool {
		return r.ID == 1
	})
	if !ok {
		t.Error("wrong")
	}
}

func TestAllBy(t *testing.T) {
	arr := []TestModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}
	ok := AllBy(arr, nil)
	if ok {
		t.Error("wrong")
		return
	}

	ok = AllBy(arr, map[string]interface{}{
		"name": "a",
	})
	if ok {
		t.Error("wrong")
		return
	}

	ok = AllBy(arr, map[string]interface{}{
		"id": 1,
	})
	if !ok {
		t.Error("wrong")
	}
}

func TestChain_All_False(t *testing.T) {
	ok := Chain([]TestModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}).All(func(r TestModel, _ int) bool {
		return r.ID != 1
	})
	if ok {
		t.Error("wrong")
	}
}

func TestChain_All_True(t *testing.T) {
	ok := Chain([]TestModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}).All(func(r TestModel, _ int) bool {
		return r.ID == 1
	})
	if !ok {
		t.Error("wrong")
	}
}

func TestChain_AllBy_False(t *testing.T) {
	ok := Chain([]TestModel{
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

func TestChain_AllBy_True(t *testing.T) {
	ok := Chain([]TestModel{
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
