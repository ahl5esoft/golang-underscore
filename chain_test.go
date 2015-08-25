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
		UnderscoreModel{ "d", 41 },
		UnderscoreModel{ "d", 4 },
		UnderscoreModel{ "c", 31 },
		UnderscoreModel{ "c", 3 },
		UnderscoreModel{ "b", 21 },
		UnderscoreModel{ "b", 2 },
		UnderscoreModel{ "a", 11 },
		UnderscoreModel{ "a", 1 },
	}
)

func TestChainGroup(t *testing.T) {
	v, err := Chain(arr).Group(func (item interface{}, _ interface{}) (interface{}, error) {
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
	v, err := Chain(arr).Index(func (item interface{}, _ interface{}) (interface{}, error) {
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

func TestChainMap(t *testing.T) {
	v, _ := Chain([]int{ 1, 2, 3 }).Map(func (item interface{}, _ interface{}) interface{} {
		return item.(int) + 10
	}).Value()

	res, ok := v.([]interface{})
	if !(ok && len(res) == 3) {
		t.Error("Chain.Map error")
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
	if !(ok && s == "d") {
		t.Error("Chain.Pluck: value error")
		return
	}

	t.Log(res)
}

func TestChainReduce(t *testing.T) {
	v, err := Chain(arr).Reduce(func (memo, value, _ interface{}) interface{} {
		dict := memo.(map[string]int)
		m := value.(UnderscoreModel)
		dict[m.Id] = m.Age
		return dict
	}, make(map[string]int)).Value()
	if err != nil {
		t.Error(err)
	}

	dict, ok := v.(map[string]int)
	if !(ok && len(dict) == 4) {
		t.Error("Chain.Reduce: vlaue error")
	}

	t.Log(dict)
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

func TestChainSortBy(t *testing.T) {
	v, err := Chain(arr).SortBy("Age").Value()
	if err != nil {
		t.Error(err)
		return
	}

	res, ok := v.([]interface{})
	if !(ok && len(res) == 8) {
		t.Error("Chain.SortBy: type error")
		return
	}

	m, ok := res[0].(UnderscoreModel)
	if !(ok && m.Id == "a") {
		t.Error("Chain.SortBy: value error")
		return
	}

	t.Log(res)
}

func TestChainUniq(t *testing.T) {
	v, _ := Chain([]int{ 1, 2, 3, 1, 4 }).Uniq().Value()

	res, ok := v.([]interface{})
	if !(ok && len(res) == 4) {
		t.Error("Chain.Uniq error")
	}
}

func TestChainUniqBy(t *testing.T) {
	v, _ := Chain([]int{ 1, 2, 3, 1, 4 }).UniqBy(func (item interface{}, _ int) interface{} {
		return item.(int) % 2
	}).Value()

	res, ok := v.([]interface{})
	if !(ok && len(res) == 2) {
		t.Error("Chain.UniqBy error")
	}
}