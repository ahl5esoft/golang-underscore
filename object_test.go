package underscore

import (
	"testing"
)

func Test_Object(t *testing.T) {
	arr := []interface{}{
		[]interface{}{"a", 1},
		[]interface{}{"b", 2},
	}
	dic, ok := Object(arr).(map[string]int)
	if !(ok && len(dic) == 2) {
		t.Error("wrong", ok, dic, len(dic))
		return
	}

	if v1, ok := dic["a"]; !(ok && v1 == 1) {
		t.Error("key 1")
		return
	}

	if v1, ok := dic["b"]; !(ok && v1 == 2) {
		t.Error("key 2")
		return
	}
}

const COUNT = 10000

func Test_objectAsParallel(t *testing.T) {
	arr := [COUNT][]int{}
	for i := 0; i < COUNT; i++ {
		arr[i] = []int{i, i}
	}
	finalArr, ok := objectAsParallel(arr).(map[int]int)
	if !(ok && len(finalArr) == COUNT) {
		t.Error("err", ok, len(finalArr))
	}
}

func Test_Chain_Object(t *testing.T) {
	arr := [COUNT][]int{}
	for i := 0; i < COUNT; i++ {
		arr[i] = []int{i, i}
	}

	finalArr, ok := Chain(arr).AsParallel().Object().Value().(map[int]int)
	if !(ok && len(finalArr) == COUNT) {
		t.Error("err", ok, len(finalArr))
	}
}
