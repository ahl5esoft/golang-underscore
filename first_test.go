package underscore

import "testing"

func Benchmark_First(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var dst int
		Range(1, benchmarkSize, 1).First().Value(&dst)
	}
}

func Benchmark_First_New(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var dst int
		Range2(1, benchmarkSize, 1).First().Value(&dst)
	}
}

func Test_First(t *testing.T) {
	var dst int
	Chain2([]int{1, 2, 3}).First().Value(&dst)
	if dst != 1 {
		t.Error("wrong")
	}
}

func Test_First_Twice(t *testing.T) {
	var dst int
	Chain2([][]int{
		[]int{1, 3, 5, 7},
		[]int{2, 4, 6, 8},
	}).First().First().Value(&dst)
	if dst != 1 {
		t.Error("wrong")
	}
}

func Test_Chain_First(t *testing.T) {
	arr := []int{1, 2, 3}
	var item int
	Chain(arr).First().Value(&item)
	if item != 1 {
		t.Error("wrong")
	}
}
