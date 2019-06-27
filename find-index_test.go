package underscore

import "testing"

func Benchmark_FindIndex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Range(1, benchmarkSize, 1).FindIndex(func(r, _ int) bool {
			return r == 200
		})
	}
}

func Test_FindIndex(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	index := Chain(src).FindIndex(func(r testModel, _ int) bool {
		return r.Name == src[1].Name
	})
	if index != 1 {
		t.Error("wrong")
	}
}

func Test_FindIndexBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	index := Chain(src).FindIndexBy(map[string]interface{}{
		"id": 1,
	})
	if index == -1 || src[index].ID != 1 {
		t.Error("wrong")
	}
}
