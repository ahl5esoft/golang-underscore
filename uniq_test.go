package underscore

import (
	"testing"
)

func TestUniq(t *testing.T) {
	v, _ := Uniq([]int{ 1, 2, 1, 2, 1, 3 }, nil)
	res, ok := v.([]int)
	if !(ok && len(res) == 3) {
		t.Error("wrong")
	}
}

func TestChain_Uniq(t *testing.T) {
	v, _ := Chain([]int{ 1, 2, 1, 4, 1, 3 }).Uniq(func (n, _ int) (int, error) {
		return n % 2, nil
	}).Value()
	res, ok := v.([]int)
	if !(ok && len(res) == 2) {
		t.Error("wrong")
	}
}

func TestUniqBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "one" },
		TestModel{ 3, "one" },
	}
	v, _ := UniqBy(arr, "Name")
	res, ok := v.([]TestModel)
	if !(ok && len(res) == 1) {
		t.Error("wrong")
	}
}

func TestChain_UniqBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ 1, "one" },
		TestModel{ 2, "one" },
		TestModel{ 3, "one" },
	}
	v, _ := Chain(arr).UniqBy("Name").Value()
	res, ok := v.([]TestModel)
	if !(ok && len(res) == 1) {
		t.Error("wrong")
	}
}