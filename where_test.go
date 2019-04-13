package underscore

import (
	"testing"
)

func Test_Where(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
		{ID: 4, Name: "three"},
	}
	res := Where(arr, func(r testModel, i int) bool {
		return r.ID%2 == 0
	}).([]testModel)
	if !(len(res) == 2 && res[0].ID == 2 && res[1].ID == 4) {
		t.Error(res)
	}
}

func Test_WhereBy(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "one"},
		{ID: 3, Name: "three"},
		{ID: 4, Name: "three"},
	}
	res := WhereBy(arr, map[string]interface{}{
		"Name": "one",
	}).([]testModel)
	if !(len(res) == 2 && res[0] == arr[0] && res[1] == arr[1]) {
		t.Error(res)
	}
}

func Test_Chain_Where(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "one"},
		{ID: 3, Name: "three"},
		{ID: 4, Name: "three"},
	}
	res := make([]testModel, 0)
	Chain(arr).Where(func(r testModel, i int) bool {
		return r.ID%2 == 0
	}).Value(&res)
	if !(len(res) == 2 && res[0].ID == 2 && res[1].ID == 4) {
		t.Error("wrong result")
	}
}

func Test_Chain_WhereBy(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "one"},
		{ID: 3, Name: "three"},
		{ID: 4, Name: "three"},
	}
	res := make([]testModel, 0)
	Chain(arr).WhereBy(map[string]interface{}{
		"Name": "one",
	}).Value(&res)
	if !(len(res) == 2 && res[0] == arr[0] && res[1] == arr[1]) {
		t.Error("wrong result")
	}
}
