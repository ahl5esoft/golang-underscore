package underscore

import "testing"

func Benchmark_Take(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dst := make([]int, 0)
		Range(1, benchmarkSize, 1).Take(200).Value(&dst)
	}
}

func Benchmark_Take_New(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dst := make([]int, 0)
		Range2(1, benchmarkSize, 1).Take(200).Value(&dst)
	}
}

func Test_Take(t *testing.T) {
	src := []int{1, 2, 3}
	dst := make([]int, 0)
	Chain2(src).Take(1).Value(&dst)
	if len(dst) != 1 || dst[0] != 1 {
		t.Fatal(dst)
	}
}
