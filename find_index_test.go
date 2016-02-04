package underscore

import (
	"testing"
)

func TestFindIndex(t *testing.T) {
	arr := []int{ 1, 2, 3, 4, 5, 6 }
	i, _ := FindIndex(arr, func (n, i int) (bool, error) {
		return n == 3, nil
	})
	if i != 2 {
		t.Error("wrong")
	}
}