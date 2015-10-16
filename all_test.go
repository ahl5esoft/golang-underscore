package underscore

import (
	"testing"
)

func TestAll(t *testing.T) {
	arr := []int{ 2, 4 }
	res, _ := All(arr, func (n, _ interface{}) (bool, error) {
		return n.(int) % 2 == 0, nil	
	})
	if !res {
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