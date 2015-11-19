package underscore

import (
	"errors"
	"testing"
)

func TestEach(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	err := Each(arr, func (n, i int) error {
		if n != arr[i] {
			return errors.New("each")
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}