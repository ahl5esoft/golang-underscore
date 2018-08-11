package underscore

import (
	"testing"
)

func Test_Object(t *testing.T) {
	arr := []interface{}{
		[]interface{}{"a", 1},
		[]interface{}{"b", 2},
	}
	dic, ok := Object(arr).(map[string]int)
	if !(ok && len(dic) == 2) {
		t.Error("wrong", ok, dic, len(dic))
		return
	}

	if v1, ok := dic["a"]; !(ok && v1 == 1) {
		t.Error("key 1")
		return
	}

	if v1, ok := dic["b"]; !(ok && v1 == 2) {
		t.Error("key 2")
		return
	}
}
