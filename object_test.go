package underscore

import (
	"testing"
)

func Test_Object(t *testing.T) {
	arr := []interface{}{
		[]interface{}{"a", 1},
		[]interface{}{"b", 2},
	}
	dic := Object(arr).(map[string]int)
	if len(dic) != 2 {
		t.Fatal(dic)
	}

	if v, ok := dic["a"]; !(ok && v == 1) {
		t.Fatal("key 1")
	}

	if v, ok := dic["b"]; !(ok && v == 2) {
		t.Error("key 2")
	}
}

func Test_objectAsParallel(t *testing.T) {
	arr := [500][]int{}
	for i := 0; i < 500; i++ {
		arr[i] = []int{i, i}
	}

	dic := objectAsParallel(arr).(map[int]int)
	if len(dic) != 500 {
		t.Error("err")
	}
}

func Test_Chain_AsParallel_Object(t *testing.T) {
	arr := [500][]int{}
	for i := 0; i < 500; i++ {
		arr[i] = []int{i, i}
	}

	res := make(map[int]int)
	Chain(arr).AsParallel().Object().Value(&res)
	if len(res) != 500 {
		t.Error(len(res))
	}
}

func Test_Chain_Object(t *testing.T) {
	arr := [500][]int{}
	for i := 0; i < 500; i++ {
		arr[i] = []int{i, i}
	}

	res := make(map[int]int)
	Chain(arr).AsParallel().Object().Value(&res)
	if len(res) != 500 {
		t.Error(len(res))
	}
}
