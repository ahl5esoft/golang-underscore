package underscore

import (
	"testing"
	"time"
)

func TestEach(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	Each(arr, func (n, i int) {
		if n != arr[i] {
			t.Error("wrong")
		}
	})
}

func TestChain_Each(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	Chain(arr).Each(func (n, i int) {
		if n != arr[i] {
			t.Error("wrong")
		}
	})
}

func TestChain_Parallel_Each(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	beginUnix := time.Now().Unix()
	Chain(arr).AsParallel().Each(func (n, i int) {
		time.Sleep(time.Second)
	}).Value()
	endUnix := time.Now().Unix()
	if int(endUnix - beginUnix) > len(arr) {
		t.Error("wrong")
	}
}