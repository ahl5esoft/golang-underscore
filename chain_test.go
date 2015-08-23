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

func TestChainPluck(t *testing.T) {
	v, err := Chain(arr).Pluck("Id").Value()
	if err != nil {
		t.Error(err)
	}

	res, ok := v.([]interface{})
	if !(ok && len(res) == 8) {
		t.Error("Chain.Pluck: type error")
		return
	}

	s, ok := res[0].(string)
	if !(ok && s == "a") {
		t.Error("Chain.Pluck: value error")
		return
	}

	t.Log(res)
}

func TestChainSize(t *testing.T) {
	v, err := Chain(arr).Size().Value()
	if err != nil {
		t.Error(err)
		return
	}

	length, ok := v.(int)
	if !(ok && length == 8) {
		t.Error("Chain.Count error")
		return
	}
}

func TestChainMap(t *testing.T) {
	v, _ := Chain([]int{ 1, 2, 3 }).Map(func (item interface{}) interface{} {
		return item.(int) + 10
	}).Value()

	res, ok := v.([]interface{})
	if !(ok && len(res) == 3) {
		t.Error("Chain.Map error")
	}
}

func TestChainUniq(t *testing.T) {
	v, _ := Chain([]int{ 1, 2, 3, 1, 4 }).Uniq().Value()

	res, ok := v.([]interface{})
	if !(ok && len(res) == 4) {
		t.Error("Chain.Uniq error")
	}
}

func TestChainUniqBy(t *testing.T) {
	v, _ := Chain([]int{ 1, 2, 3, 1, 4 }).UniqBy(func (item interface{}) interface{} {
		return item.(int) % 2
	}).Value()

	res, ok := v.([]interface{})
	if !(ok && len(res) == 2) {
		t.Error("Chain.UniqBy error")
	}
}