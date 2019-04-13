package underscore

import (
	"testing"
)

func Test_FindIndex(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}
	i := FindIndex(arr, func(r testModel, _ int) bool {
		return r.Name == arr[1].Name
	})
	if i != 1 {
		t.Error("wrong")
	}
}

func Test_FindIndexBy(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	i := FindIndexBy(arr, map[string]interface{}{
		"id": 1,
	})
	if i != 0 {
		t.Error("wrong")
	}
}

func Test_Chain_FindIndex(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	index := Chain(arr).FindIndex(func(r testModel, _ int) bool {
		return r.Name == arr[1].Name
	})
	if index != 1 {
		t.Error("wrong")
	}
}

func Test_Chain_FindIndexBy(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	index := Chain(arr).FindIndexBy(map[string]interface{}{
		"id": 1,
	})
	if index == -1 || arr[index].ID != 1 {
		t.Error("wrong")
	}
}
