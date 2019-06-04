package underscore

import "testing"

func Test_Filter(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "one"},
		{ID: 3, Name: "three"},
		{ID: 4, Name: "three"},
	}
	dst := make([]testModel, 0)
	Chain2(src).Filter(func(r testModel, _ int) bool {
		return r.ID%2 == 0
	}).Value(&dst)
	if !(len(dst) == 2 && dst[0] == src[1] && dst[1] == src[3]) {
		t.Error(dst)
	}
}

func Test_FilterBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "one"},
		{ID: 3, Name: "three"},
		{ID: 4, Name: "three"},
	}
	dst := make([]testModel, 0)
	Chain2(src).FilterBy(map[string]interface{}{
		"Name": "one",
	}).Value(&dst)
	if !(len(dst) == 2 && dst[0] == src[0] && dst[1] == src[1]) {
		t.Error("wrong result")
	}
}
