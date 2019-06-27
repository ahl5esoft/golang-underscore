package underscore

import "testing"

func Benchmark_Find(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var dst int
		Range(1, benchmarkSize, 1).Find(func(r, _ int) bool {
			return r > 0
		}).Value(&dst)
	}
}

func Test_Chain_Find(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	res := testModel{}
	Chain(arr).Find(func(r testModel, _ int) bool {
		return r.ID == 1
	}).Value(&res)
	if res != arr[0] {
		t.Error("wrong")
	}
}

func Test_Chain_FindBy(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	res := testModel{}
	Chain(arr).FindBy(map[string]interface{}{
		"id": 2,
	}).Value(&res)
	if res != arr[1] {
		t.Error("wrong")
	}
}

func Test_Find(t *testing.T) {
	var dst int
	Chain([]int{1, 2, 3}).Find(func(r, _ int) bool {
		return r == 2
	}).Value(&dst)
	if dst != 2 {
		t.Error("wrong")
	}
}

func Test_FindBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	var dst testModel
	Chain(src).FindBy(map[string]interface{}{
		"id": 2,
	}).Value(&dst)
	if dst != src[1] {
		t.Error("wrong")
	}
}
