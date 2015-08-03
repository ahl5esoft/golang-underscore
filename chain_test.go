package underscore

import (
	"testing"
)

func TestChain(t *testing.T) {
	res, ok := Chain(arr).Index(func (item interface{}) interface{} {
		return item.(UnderscoreModel).Id
	}).IndexBy("Id").Value().(map[interface{}]interface{})
	if !ok {
		t.Error("wrong type")
	}

	m, ok := res["a"].(UnderscoreModel)
	if !(ok && m.Id == "a") {
		t.Error("wrong model")
	}

	t.Log(m)
}