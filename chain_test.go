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
	v, err := Chain(arr).Group(func (item interface{}) (interface{}, error) {
		return item.(UnderscoreModel).Id, nil
	}).Value()
	if err != nil {
		t.Error(err)
		return
	}

	res, ok := v.(map[interface{}][]interface{})
	if !ok {
		t.Error("Chain.Group: type error")
		return
	}

	group, ok := res["a"]
	if !(ok && len(group) == 2) {
		t.Error("Chain.Group: value error")
		return
	}
}

func TestChainGroupBy(t *testing.T) {
	_, err := Chain(arr).GroupBy("id").Value()
	if err == nil {
		t.Error("GroupBy has BUG!")
	}
}

func TestChainIndex(t *testing.T) {
	v, err := Chain(arr).Index(func (item interface{}) (interface{}, error) {
		return item.(UnderscoreModel).Id, nil
	}).Value()
	if err != nil {
		t.Error(err)
		return
	}

	res, ok := v.(map[interface{}]interface{})
	if !ok {
		t.Error("Chain.Index: type error")
		return
	}

	m, ok := res["a"].(UnderscoreModel)
	if !(ok && m.Id == "a") {
		t.Error("Chain.Index: value error")
		return
	}
}

func TestChainIndexBy(t *testing.T) {
	_, err := Chain(arr).IndexBy("id").Value()
	if err == nil {
		t.Error("IndexBy has BUG!")
	}
}

func TestChainCount(t *testing.T) {
	v, err := Chain(arr).Count().Value()
	if err != nil {
		t.Error(err)
		return
	}

	count, ok := v.(int)
	if !(ok && count == 8) {
		t.Error("Chain.Count error")
		return
	}
}