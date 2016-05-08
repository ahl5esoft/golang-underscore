package underscore

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	m := TestModel{ 1, "one" }
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
		"id": m.Id,
		"name": "a",
	})
	if ok {
		t.Error("wrong")
		return
	}

	ok = IsMatch(m, map[string]interface{}{
		"id": m.Id,
		"name": m.Name,
	})
	if !ok {
		t.Error("wrong")
	}
}