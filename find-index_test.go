package underscore

import (
	"testing"
)

func TestFindIndex(t *testing.T) {
	arr := []int{ 1, 2, 3, 4, 5, 6 }
	i := FindIndex(arr, func (n, _ int) bool {
		return n == 3
	})
	if i == -1 || arr[i] != 3 {
		t.Error("wrong")
	}
}

func TestChain_FindIndex(t *testing.T) {
	arr := []int{ 1, 2, 3, 4, 5, 6 }
	res := Chain(arr).FindIndex(func (n, _ int) bool {
		return n == 3
	}).Value()
	index := res.(int)
	if index == -1 || arr[index] != 3 {
		t.Error("wrong")
	}
}

func TestFindIndexBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	i := FindIndexBy(arr, map[string]interface{}{
		"id": 1,
	})
	if i == -1 || arr[i].Id != 1 {
		t.Error("wrong")
	}
}

func TestChain_FindIndexBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	res := Chain(arr).FindIndexBy(map[string]interface{}{
		"id": 1,
	}).Value()
	i := res.(int)
	if i == -1 || arr[i].Id != 1 {
		t.Error("wrong")
	}
}