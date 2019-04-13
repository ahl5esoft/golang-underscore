package underscore

import (
	"testing"
)

func Test_IsMatch(t *testing.T) {
	m := testModel{ID: 1, Name: "one"}
	ok := IsMatch(nil, nil)
	if ok {
		t.Error("wrong")
		return
	}

	ok = IsMatch(m, nil)
	if ok {
		t.Error("wrong")
		return
	}

	ok = IsMatch(m, map[string]interface{}{
		"id":   m.ID,
		"name": "a",
	})
	if ok {
		t.Error("wrong")
		return
	}

	ok = IsMatch(m, map[string]interface{}{
		"id":   m.ID,
		"name": m.Name,
	})
	if !ok {
		t.Error("wrong")
	}
}
