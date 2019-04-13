package underscore

import (
	"testing"
	"time"
)

func Test_Each(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}
	Each(arr, func(r testModel, i int) {
		if !(r.ID == arr[i].ID && r.Name == arr[i].Name) {
			t.Error("wrong")
		}
	})
}

func Test_Chain_Each(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}
	Chain(arr).Each(func(r testModel, i int) {
		if !(r.ID == arr[i].ID && r.Name == arr[i].Name) {
			t.Error("wrong")
		}
	})
}

func Test_Chain_Parallel_Each(t *testing.T) {
	arr := []int{1, 2, 3}
	beginUnix := time.Now().Unix()
	Chain(arr).AsParallel().Each(func(n, i int) {
		time.Sleep(time.Second)
	})
	endUnix := time.Now().Unix()
	if int(endUnix-beginUnix) > len(arr) {
		t.Error("wrong")
	}
}
