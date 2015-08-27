package underscore

import (
	"testing"
)

func TestMap(t *testing.T) {
	arr := []string{ "a", "b", "c" }
	res, _ := Map(arr, func (item interface{}, _ interface{}) (interface{}, error) {
		return item.(string) + "-", nil
	})
	if len(res) != len(arr) {
		t.Error("Map has diff len")
	}

	if res[0].(string) != "a-" {
		t.Error("Map BUG")
	}
}