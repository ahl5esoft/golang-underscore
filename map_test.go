package underscore

import (
	"testing"
)

func TestMap(t *testing.T) {
	arr := []string{ "a", "b", "c" }
	res, _ := Map(arr, func (item, _ interface{}) (interface{}, error) {
		return item.(string) + "-", nil
	})
	if !(len(res) == len(arr) && res[0].(string) == "a-") {
		t.Error("wrong")
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