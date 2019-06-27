package underscore

import "testing"

func Benchmark_Group(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dst := make([]int, 0)
		Range(1, benchmarkSize, 1).Group(func(n, _ int) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		}).Value(&dst)
	}
}

func Benchmark_Group_New(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dst := make([]int, 0)
		Range(1, benchmarkSize, 1).Group(func(n, _ int) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		}).Value(&dst)
	}
}

func Test_Group(t *testing.T) {
	dst := make(map[string][]int)
	Chain([]int{1, 2, 3, 4, 5}).Group(func(n, _ int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	}).Value(&dst)
	if len(dst["even"]) != 2 {
		t.Error("wrong")
	}
}

func Test_GroupBy(t *testing.T) {
	dst := make(map[string][]testModel)
	Chain([]testModel{
		{ID: 1, Name: "a"},
		{ID: 2, Name: "a"},
		{ID: 3, Name: "b"},
		{ID: 4, Name: "b"},
	}).GroupBy("Name").Value(&dst)
	if len(dst) != 2 {
		t.Error("wrong")
	}
}
