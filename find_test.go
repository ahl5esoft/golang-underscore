package underscore

import (
	"testing"
)

func Test_Find(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	item := Find(arr, func(r testModel, _ int) bool {
		return r.ID == 1
	})
	if item != arr[0] {
		t.Error("wrong")
	}
}

func Test_FindBy(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	item := FindBy(arr, map[string]interface{}{
		"id": 2,
	})
	if item != arr[1] {
		t.Error("wrong")
	}
}

func Test_Chain_Find(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	res := testModel{}
	Chain(arr).Find(func(r testModel, _ int) bool {
		return r.ID == 1
	}).Value(&res)
	if res != arr[0] {
		t.Error("wrong")
	}
}

func Test_Chain_FindBy(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	res := testModel{}
	Chain(arr).FindBy(map[string]interface{}{
		"id": 2,
	}).Value(&res)
	if res != arr[1] {
		t.Error("wrong")
	}
}
