package underscore

import (
	"testing"
)

type mapManytestModel struct {
	Str   string
	Slice []string
	Array [2]int
}

func Test_MapMany_NotSliceOrArray(t *testing.T) {
	defer func() {
		if rv := recover(); rv == nil {
			t.Error("wrong")
		}
	}()

	src := []string{"a", "b", "c"}
	MapMany(src, func(v string, _ int) string {
		return v
	})
}

func Test_Chain_MapMany_NotSliceOrArray(t *testing.T) {
	defer func() {
		if rv := recover(); rv == nil {
			t.Error("wrong")
		}
	}()

	src := []string{"a", "b", "c"}
	Chain(src).MapMany(func(v string, _ int) string {
		return v
	})
}

func Test_MapManyBy_PropertyNotSliceOrArray(t *testing.T) {
	defer func() {
		if rv := recover(); rv == nil {
			t.Error("wrong")
		}
	}()

	src := []mapManytestModel{
		{"", nil, [2]int{}},
		{"", nil, [2]int{}},
		{"", nil, [2]int{}},
	}
	MapManyBy(src, "Str")
}

func Test_Chain_MapManyBy_PropertyNotSliceOrArray(t *testing.T) {
	defer func() {
		if rv := recover(); rv == nil {
			t.Error("wrong")
		}
	}()

	src := []mapManytestModel{
		{"", nil, [2]int{}},
		{"", nil, [2]int{}},
		{"", nil, [2]int{}},
	}
	Chain(src).MapManyBy("Str")
}

func Test_MapMany_Slice(t *testing.T) {
	src := []int{1, 2}
	res := MapMany(src, func(r, _ int) []int {
		var temp []int
		Range(0, r, 1).Map(func(_, _ int) int {
			return r
		}).Value(&temp)
		return temp
	}).([]int)
	if !(len(res) == 3 && res[0] == 1 && res[1] == 2 && res[2] == 2) {
		t.Error(res)
	}
}

func Test_Chain_MapMany_Slice(t *testing.T) {
	src := []int{1, 2}
	var res []int
	Chain(src).MapMany(func(r, _ int) []int {
		var temp []int
		Range(0, r, 1).Map(func(_, _ int) int {
			return r
		}).Value(&temp)
		return temp
	}).Value(&res)
	if !(len(res) == 3 && res[0] == 1 && res[1] == 2 && res[2] == 2) {
		t.Error(res)
	}
}

func Test_MapManyBy_Slice(t *testing.T) {
	src := []mapManytestModel{
		{"", []string{"a", "b"}, [2]int{}},
		{"", []string{"c", "d"}, [2]int{}},
	}
	res := MapManyBy(src, "Slice").([]string)
	if !(len(res) == 4 && res[0] == "a" && res[1] == "b" && res[2] == "c" && res[3] == "d") {
		t.Error(res)
	}
}

func Test_Chain_MapManyBy_Slice(t *testing.T) {
	src := []mapManytestModel{
		{"", []string{"a", "b"}, [2]int{}},
		{"", []string{"c", "d"}, [2]int{}},
	}
	var res []string
	Chain(src).MapManyBy("Slice").Value(&res)
	if !(len(res) == 4 && res[0] == "a" && res[1] == "b" && res[2] == "c" && res[3] == "d") {
		t.Error(res)
	}
}

func Test_MapMany_Array(t *testing.T) {
	src := [2]int{1, 2}
	res := MapMany(src, func(r, _ int) []int {
		var temp []int
		Range(0, r, 1).Map(func(_, _ int) int {
			return r
		}).Value(&temp)
		return temp
	}).([]int)
	if !(len(res) == 3 && res[0] == 1 && res[1] == 2 && res[2] == 2) {
		t.Error(res)
	}
}

func Test_Chain_MapMany_Array(t *testing.T) {
	src := [2]int{1, 2}
	var res []int
	Chain(src).MapMany(func(r, _ int) []int {
		var temp []int
		Range(0, r, 1).Map(func(_, _ int) int {
			return r
		}).Value(&temp)
		return temp
	}).Value(&res)
	if !(len(res) == 3 && res[0] == 1 && res[1] == 2 && res[2] == 2) {
		t.Error(res)
	}
}

func Test_MapManyBy_Array(t *testing.T) {
	src := []mapManytestModel{
		{"", nil, [2]int{1, 2}},
		{"", nil, [2]int{3, 4}},
	}
	res := MapManyBy(src, "Array").([]int)
	if !(len(res) == 4 && res[0] == 1 && res[1] == 2 && res[2] == 3 && res[3] == 4) {
		t.Error(res)
	}
}

func Test_Chain_MapManyBy_Array(t *testing.T) {
	src := []mapManytestModel{
		{"", nil, [2]int{1, 2}},
		{"", nil, [2]int{3, 4}},
	}
	var res []int
	Chain(src).MapManyBy("Array").Value(&res)
	if !(len(res) == 4 && res[0] == 1 && res[1] == 2 && res[2] == 3 && res[3] == 4) {
		t.Error(res)
	}
}
