package underscore

import (
	"testing"
)

func TestIndex(t *testing.T) {
	res := Index([]string{ "a", "b" }, func (item interface{}) interface{} {
		return item
	})
	str, ok := res["a"].(string)
	if !(ok && str == "a") {
		t.Error("index error")
	}
}