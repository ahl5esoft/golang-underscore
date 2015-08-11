package underscore

import (
	"testing"
)

type UnderscoreModel struct {
	Id string
	Age int
}

var (
	arr = []UnderscoreModel{
		UnderscoreModel{ "a", 1 },
		UnderscoreModel{ "a", 11 },
		UnderscoreModel{ "b", 2 },
		UnderscoreModel{ "b", 21 },
		UnderscoreModel{ "c", 3 },
		UnderscoreModel{ "c", 31 },
		UnderscoreModel{ "d", 4 },
		UnderscoreModel{ "d", 41 },
	}
)

func TestChainGroup(t *testing.T) {
	res, ok := Chain(arr).Group(func (item interface{}) interface{} {
		return item.(UnderscoreModel).Id
	}).Value().(map[interface{}][]interface{})
	if !ok {
		t.Error("Chain.Group: type error")
	}

	group, ok := res["a"]
	if !(ok && len(group) == 2) {
		t.Error("Chain.Group: value error")
	}
}

func TestChainIndex(t *testing.T) {
	res, ok := Chain(arr).Index(func (item interface{}) interface{} {
		return item.(UnderscoreModel).Id
	}).Value().(map[interface{}]interface{})
	if !ok {
		t.Error("Chain.Index: type error")
	}

	m, ok := res["a"].(UnderscoreModel)
	if !(ok && m.Id == "a") {
		t.Error("Chain.Index: value error")
	}
}

func TestChainIndexBy(t *testing.T) {
	res, ok := Chain(arr).IndexBy("Id").Value().(map[interface{}]interface{})
	if !ok {
		t.Error("chain.IndexBy: type error")
	}
	
	m, ok := res["a"].(UnderscoreModel)
	if !(ok && m.Id == "a") {
		t.Error("Chain.Index: value error")
	}
}

func TestChainCount(t *testing.T) {
	count, ok := Chain(arr).Count().Value().(int)
	if !(ok && count == 8) {
		t.Error("Chain.Count error")
	}
}