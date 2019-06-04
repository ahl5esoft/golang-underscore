package underscore

import (
	"testing"
)

func Test_Each(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 1, Name: "two"},
		{ID: 1, Name: "three"},
	}
	Chain2(arr).Each(func(r testModel, i int) {
		if !(r.ID == arr[i].ID && r.Name == arr[i].Name) {
			t.Error("wrong")
		}
	})
}
