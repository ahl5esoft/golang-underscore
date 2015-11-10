package underscore

import (
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	arr := []string{ "11", "12", "13" }
	v, err := Map(arr, func (s string, _ int) (int, error) {
		return strconv.Atoi(s)
	})
	if err != nil {
		t.Error(err)
		return
	}

	res, ok := v.([]int)
	if !(ok && len(res) == len(arr)) {
		t.Error("wrong type")
		return
	}

	for i, s := range arr {
		n, _ := strconv.Atoi(s)
		if n != res[i] {
			t.Error("wrong value")
			return;
		}
	}
}

func TestChain_Map(t *testing.T) {
	arr := []string{ "a", "b", "c" }
	v, _ := Chain(arr).Map(func (item, _ interface{}) (interface{}, error) {
		return item.(string) + "-", nil
	}).Value()

	res, ok := v.([]interface{})
	if !(ok && len(res) == len(arr) && res[0].(string) == "a-") {
		t.Error("wrong")
	}
}