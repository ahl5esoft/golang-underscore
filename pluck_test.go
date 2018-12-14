package underscore

import (
	"testing"
)

func TestPluck(t *testing.T) {
	arr := []TestModel{
		TestModel{1, "one"},
		TestModel{2, "two"},
		TestModel{3, "three"},
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
		TestModel{1, "one"},
		TestModel{2, "two"},
		TestModel{3, "three"},
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

type mainModel struct {
	nested

	Name string
}

type nested struct {
	ID string
}

func TestPluck_Nested(t *testing.T) {
	m := mainModel{
		nested: nested{
			ID: "nested",
		},
		Name: "name",
	}
	v := Pluck([]mainModel{m}, "id")
	ids, ok := v.([]string)
	if !(ok && len(ids) == 1 && ids[0] == m.ID) {
		t.Error(v)
	}
}
