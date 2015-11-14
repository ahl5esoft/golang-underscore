package underscore

import (
	"testing"
)

func TestAny(t *testing.T) {
	arr := []int{ 1, 3 }
	res, _ := Any(arr, func (n, _ int) (bool, error) {
		return n % 2 == 0, nil	
	})
	if res {
		t.Error("wrong result")
	}
}

func TestChain_Any(t *testing.T) {
	arr := []int{ 1, 3 }
	v, _ := Chain(arr).Any(func (n, _ int) (bool, error) {
		return n % 2 == 0, nil	
	}).Value()
	res, ok := v.(bool)
	if !ok || res {
		t.Error("wrong result")
	}
}

func TestAnyBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	res, err := AnyBy(arr, map[string]interface{}{
		"Id": 2,
		"Name": "two",
	})
	if err != nil {
		t.Error(err)
	}

	if !res {
		t.Error("wrong result")
	}
}

func TestChain_AnyBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	v, err := Chain(arr).AnyBy(map[string]interface{}{
		"Id": 2,
		"Name": "two",
	}).Value()
	if err != nil {
		t.Error(err)
	}

	res, ok := v.(bool)
	if !(ok && res) {
		t.Error("wrong result")
	}
}