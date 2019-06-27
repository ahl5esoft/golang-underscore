package underscore

import "testing"

func Test_Index(t *testing.T) {
	src := []string{"a", "b"}
	dst := make(map[string]string)
	Chain(src).Index(func(item string, _ int) string {
		return item
	}).Value(&dst)
	if len(dst) != 2 || dst["a"] != "a" || dst["b"] != "b" {
		t.Error(dst)
	}
}

func Test_IndexBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "a"},
		{ID: 2, Name: "a"},
		{ID: 3, Name: "b"},
		{ID: 4, Name: "b"},
	}
	dst := make(map[string]testModel)
	Chain(src).IndexBy("name").Value(&dst)
	if len(dst) != 2 {
		t.Error(dst)
	}
}
