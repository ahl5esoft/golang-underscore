package underscore

import (
	"testing"
)

func TestPropertyRV(t *testing.T) {
	item := TestModel{ID: 1, Name: "one"}

	rv := PropertyRV("$$")(item)
	if rv != NullRv {
		t.Fatal("wrong")
	}

	getNameRV := PropertyRV("name")
	nameRV := getNameRV(item)
	if nameRV.String() != item.Name {
		t.Error("wrong")
	}
}

func TestProperty(t *testing.T) {
	item := TestModel{ID: 1, Name: "one"}

	rv := PropertyRV("$$")(item)
	if rv != NullRv {
		t.Fatal("wrong")
	}

	getName := Property("name")
	name := getName(item)
	if name.(string) != item.Name {
		t.Error("wrong")
	}
}

func TestProperty_Ptr(t *testing.T) {
	item := &TestModel{ID: 1, Name: "ptr"}

	nameGetter := Property("name")
	name := nameGetter(item)
	if name != item.Name {
		t.Error(name)
	}
}

func TestProperty_Nested(t *testing.T) {
	item := TestModel{
		TestNestedModel: TestNestedModel{
			Age: 11,
		},
	}
	age, ok := Property("age")(item).(int)
	if !(ok && age == 11) {
		t.Error("err")
	}
}
