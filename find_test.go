package underscore

import (
	"testing"
)

func TestFind(t *testing.T) {
	arr := []int{ 1, 2, 3, 4 }
	n := Find(arr, func (n, _ int) bool {
		return n % 2 == 0
	})
	if n != 2 {
		t.Error("wrong")
	}
}

func TestChain_Find(t *testing.T) {
	arr := []int{ 1, 2, 3, 4 }
	res := Chain(arr).Find(func (n, _ int) bool {
		return n % 2 == 0
	}).Value()
	if res.(int) != 2 {
		t.Error("wrong")
	}
}

func TestFindBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	res := FindBy(arr, map[string]interface{}{
		"id": 1,
	})
	if res == nil {
		t.Error("wrong")
		return
	}

	matcher := res.(TestModel)
	if !(matcher.Id == arr[0].Id && matcher.Name == arr[0].Name) {
		t.Error("wrong")
	}
}

func TestChain_FindBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	res := Chain(arr).FindBy(map[string]interface{}{
		"id": 1,
	}).Value()
	if res == nil {
		t.Error("wrong")
		return
	}

	matcher := res.(TestModel)
	if !(matcher.Id == arr[0].Id && matcher.Name == arr[0].Name) {
		t.Error("wrong")
	}
}