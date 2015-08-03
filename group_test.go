package underscore

import (
	"testing"
)

func TestGroup(t *testing.T) {
	dict := Group(arr, func (item interface{}) interface{} {
		return item.(UnderscoreModel).Id
	})
	group, ok := dict["a"]
	if !(ok && len(group) == 2) {
		t.Error("wrong")
	}
	t.Log(dict)
}