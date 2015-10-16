package underscore

import (
	"testing"
)

func TestReduce(t *testing.T) {
	v, err := Reduce([]int{ 1, 2 }, func (memo, value, _ interface{}) (interface{}, error) {
		arr := memo.([]int)
		num := value.(int)
		arr = append(arr, num)
		arr = append(arr, num + 10)
		return arr, nil
	}, make([]int, 0))	
	if err != nil {
		t.Error(err)
	}

	res, ok := v.([]int)
	if !(ok && len(res) == 4) {
		t.Error("wrong length")
	}

	if !(res[0] == 1 && res[1] == 11 && res[2] == 2 && res[3] == 12) {
		t.Error("wrong result")
	}
}

func TestChain_Reduce(t *testing.T) {
	v, err := Chain([]int{ 1, 2 }).Reduce(func (memo, value, _ interface{}) (interface{}, error) {
		arr := memo.([]int)
		num := value.(int)
		arr = append(arr, num)
		arr = append(arr, num + 10)
		return arr, nil
	}, make([]int, 0)).Value()	
	if err != nil {
		t.Error(err)
	}

	res, ok := v.([]int)
	if !(ok && len(res) == 4) {
		t.Error("wrong length")
	}

	if !(res[0] == 1 && res[1] == 11 && res[2] == 2 && res[3] == 12) {
		t.Error("wrong result")
	}
}