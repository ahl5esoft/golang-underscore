package underscore

import (
	"testing"
)

func TestSelect(t *testing.T) {
	arr := []int{ 1, 2, 3, 4 }
	res, _ := Select(arr, func (n, _ interface{}) (bool, error) {
		return n.(int) % 2 == 0, nil
	})
	if len(res) != 2 {
		t.Error("wrong length")
	}

	if !(res[0].(int) == 2 && res[1].(int) == 4) {
		t.Error("wrong result")
	}
}

func TestChain_Select(t *testing.T) {
	arr := []int{ 1, 2, 3, 4 }
	v, _ := Chain(arr).Select(func (n, _ interface{}) (bool, error) {
		return n.(int) % 2 == 0, nil
	}).Value()
	res, ok := v.([]interface{})
	if !(ok && len(res) == 2) {
		t.Error("wrong length")
	}

	
	if !(res[0].(int) == 2 && res[1].(int) == 4) {
		t.Error("wrong result")
	}
}

func TestSelectBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	res, err := SelectBy(arr, map[string]interface{}{
		"Id": 1,
	})
	if err != nil {
		t.Error(err)
	}

	if !(len(res) == 1 && res[0].(TestModel) == arr[0]) {
		t.Error("wrong result")
	}
}

func TestChain_SelectBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	v, err := Chain(arr).SelectBy(map[string]interface{}{
		"Id": 1,
	}).Value()
	if err != nil {
		t.Error(err)
	}

	res, ok := v.([]interface{})
	if !(ok && len(res) == 1 && res[0].(TestModel) == arr[0]) {
		t.Error("wrong result")
	}
}