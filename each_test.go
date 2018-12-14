package underscore

import (
	"testing"
	"time"
)

func TestEach(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 1, Name: "two"},
		TestModel{ID: 1, Name: "three"},
	}
	Each(arr, func(r TestModel, i int) {
		if !(r.ID == arr[i].ID && r.Name == arr[i].Name) {
			t.Error("wrong")
		}
	})
}

func TestChain_Each(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 1, Name: "two"},
		TestModel{ID: 1, Name: "three"},
	}
	Chain(arr).Each(func(r TestModel, i int) {
		if !(r.ID == arr[i].ID && r.Name == arr[i].Name) {
			t.Error("wrong")
		}
	})
}

func TestChain_Parallel_Each(t *testing.T) {
	arr := []int{1, 2, 3}
	beginUnix := time.Now().Unix()
	Chain(arr).AsParallel().Each(func(n, i int) {
		time.Sleep(time.Second)
	}).Value()
	endUnix := time.Now().Unix()
	if int(endUnix-beginUnix) > len(arr) {
		t.Error("wrong")
	}
}
