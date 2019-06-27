package underscore

import "testing"

func Benchmark_Any_New(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Range(1, benchmarkSize, 1).Any(func(r, _ int) bool {
			return r > 1000
		})
	}
}

func Test_Any_False(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).Any(func(r testModel, _ int) bool {
		return r.ID == 0
	})
	if ok {
		t.Error("wrong")
	}
}

func Test_Any_True(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).Any(func(r testModel, _ int) bool {
		return r.ID == 1
	})
	if !ok {
		t.Error("wrong")
	}
}

func Test_AnyBy_False(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).AnyBy(map[string]interface{}{
		"id": 0,
	})
	if ok {
		t.Error("wrong")
	}
}

func Test_AnyBy_True(t *testing.T) {
	ok := Chain([]testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}).AnyBy(map[string]interface{}{
		"name": "two",
	})
	if !ok {
		t.Error("wrong")
	}
}
