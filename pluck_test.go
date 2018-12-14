package underscore

import (
	"testing"
)

func TestPluck(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	v := Pluck(arr, "name")
	res, ok := v.([]string)
	if !(ok && len(res) == len(arr)) {
		t.Fatal("wrong length")
	}

	for i := 0; i < 3; i++ {
		if res[i] != arr[i].Name {
			t.Error("wrong result")
		}
	}
}

func TestChain_Pluck(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	v := Chain(arr).Pluck("Name").Value()
	res, ok := v.([]string)
	if !(ok && len(res) == len(arr)) {
		t.Fatal("wrong length")
	}

	for i := 0; i < 3; i++ {
		if res[i] != arr[i].Name {
			t.Error("wrong result")
		}
	}
}
