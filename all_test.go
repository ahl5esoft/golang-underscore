package underscore

import (
	"testing"
)

func TestAll(t *testing.T) {
	arr := []int{ 2, 4 }
	res, _ := All(arr, func (n, _ int) (bool, error) {
		return n % 2 == 0, nil	
	})
	if !res {
		t.Error("wrong result")
	}
}

func TestChain_All(t *testing.T) {
	arr := []int{ 2, 4 }
	res, _ := Chain(arr).All(func (n, _ int) (bool, error) {
		return n % 2 == 0, nil	
	}).Value()
	if !res.(bool) {
		t.Error("wrong result")
	}
}

func TestAllBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	res, err := AllBy(arr, map[string]interface{}{
		"Name": "a",
	})
	if err != nil {
		t.Error(err)
	}

	if res {
		t.Error("wrong result")
	}
}

func TestChain_AllBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	res, err := Chain(arr).AllBy(map[string]interface{}{
		"Name": "a",
	}).Value()
	if err != nil {
		t.Error(err)
	}

	if res.(bool) {
		t.Error("wrong result")
	}
}