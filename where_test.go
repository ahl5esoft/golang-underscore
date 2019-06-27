package underscore

import "testing"

func Benchmark_Where(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Range(0, benchmarkSize, 1).Where(func(r, _ int) bool {
			return r > 100
		}).Any(func(r, _ int) bool {
			return true
		})
	}
}

func Test_Where(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "one"},
		{ID: 3, Name: "three"},
		{ID: 4, Name: "three"},
	}
	dst := make([]testModel, 0)
	Chain(src).Where(func(r testModel, _ int) bool {
		return r.ID%2 == 0
	}).Value(&dst)
	if !(len(dst) == 2 && dst[0] == src[1] && dst[1] == src[3]) {
		t.Error(dst)
	}
}

func Test_WhereBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "one"},
		{ID: 3, Name: "three"},
		{ID: 4, Name: "three"},
	}
	dst := make([]testModel, 0)
	Chain(src).WhereBy(map[string]interface{}{
		"Name": "one",
	}).Value(&dst)
	if !(len(dst) == 2 && dst[0] == src[0] && dst[1] == src[1]) {
		t.Error("wrong result")
	}
}
