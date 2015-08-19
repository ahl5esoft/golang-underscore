package underscore

import (
	"testing"
)

func TestIndex(t *testing.T) {
	res, _ := Index([]string{ "a", "b" }, func (item interface{}) (interface{}, error) {
		return item, nil
	})
	str, ok := res["a"].(string)
	if !(ok && str == "a") {
		t.Error("index error")
	}
}