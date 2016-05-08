package underscore

import (
	"testing"
)

func TestAll(t *testing.T) {
	arr := []int{ 2, 4 }
	ok := All(arr, func (n, _ int) bool {
		return n % 2 == 0
	})
	if !ok {
		t.Error("wrong")
	}
}

func TestChain_All(t *testing.T) {
	arr := []int{ 2, 4 }
	res := Chain(arr).All(func (n, _ int) bool {
		return n % 2 == 0
	}).Value()
	if !res.(bool) {
		t.Error("wrong")
	}
}

func TestAllBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 1, "two" },
		TestModel{ 1, "three" },
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

func TestChain_AllBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	res := Chain(arr).AllBy(map[string]interface{}{
		"Name": "a",
	}).Value()
	if res.(bool) {
		t.Error("wrong")
	}
}