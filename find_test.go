package underscore

import (
	"testing"
)

func TestFind(t *testing.T) {
	arr := []int{ 1, 2, 3, 4 }
	res, _ := Find(arr, func (n, _ interface{}) (bool, error) {
		return n.(int) % 2 == 0, nil
	})
	if res == nil {
		t.Error("wrong")
	} else {
		v, ok := res.(int)
		if !(ok && v == 2) {
			t.Error("wrong")
		}
	}
}

func TestChain_Find(t *testing.T) {
	arr := []int{ 1, 2, 3, 4 }
	res, _ := Chain(arr).Find(func (n, _ interface{}) (bool, error) {
		return n.(int) % 2 == 0, nil
	}).Value()
	if res == nil {
		t.Error("wrong")
	} else {
		v, ok := res.(int)
		if !(ok && v == 2) {
			t.Error("wrong")
		}
	}
}

func TestFindBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	res, err := FindBy(arr, map[string]interface{}{
		"Id": 1,
	})
	if err != nil || res == nil {
		t.Error("wrong")
	}

	m, ok := res.(TestModel)
	if !(ok && m.Name == "one") {
		t.Error("wrong")
	}
}

func TestChain_FindBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "two" },
		TestModel{ 3, "three" },
	}
	res, err := Chain(arr).FindBy(map[string]interface{}{
		"Id": 1,
	}).Value()
	if err != nil || res == nil {
		t.Error("wrong")
	}

	m, ok := res.(TestModel)
	if !(ok && m.Name == "one") {
		t.Error("wrong")
	}
}