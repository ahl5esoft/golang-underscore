package underscore

import (
	"testing"
)

type PluckModel struct {
	Id int
	Name string
}

func TestPluck(t *testing.T) {
	arr := []PluckModel{
		PluckModel{ 1, "one" },
		PluckModel{ 2, "two" },
		PluckModel{ 3, "three" },
	}
	res, err := Pluck(arr, "Name")
	if err != nil {
		t.Error(err)
	}

	if len(res) != 3 {
		t.Error("wrong size")
	}

	for i := 0; i < 3; i++ {
		if res[i].(string) != arr[i].Name {
			t.Error("wrong result")
		}
	}
}