package underscore

import (
	"testing"
)

func TestIndex(t *testing.T) {
	res := Index(arr, func (v interface{}) interface{} {
		return v.(UnderscoreModel).Id
	})
	v, ok := res["a"].(UnderscoreModel)
	if !(ok && v.Id == "a") {
		t.Error("[array]wrong model")
	}
	t.Log(res)

	res = IndexBy(res, "Id")
	v, ok = res["b"].(UnderscoreModel)
	if !(ok && v.Id == "b") {
		t.Error("[map]wrong model")
	}
	t.Log(res)
}