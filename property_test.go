package underscore

import (
	"testing"
)

func Test_PropertyRV(t *testing.T) {
	item := testModel{ID: 1, Name: "one"}

	rv := PropertyRV("$$")(item)
	if rv != nilRV {
		t.Fatal("wrong")
	}

	getNameRV := PropertyRV("name")
	nameRV := getNameRV(item)
	if nameRV.String() != item.Name {
		t.Error("wrong")
	}
}

func Test_Property(t *testing.T) {
	item := testModel{ID: 1, Name: "one"}

	rv := PropertyRV("$$")(item)
	if rv != nilRV {
		t.Fatal("wrong")
	}

	getName := Property("name")
	name := getName(item)
	if name.(string) != item.Name {
		t.Error("wrong")
	}
}

func Test_Property_Ptr(t *testing.T) {
	item := &testModel{ID: 1, Name: "ptr"}

	nameGetter := Property("name")
	name := nameGetter(item)
	if name != item.Name {
		t.Error(name)
	}
}

func Test_Property_Nested(t *testing.T) {
	item := testModel{
		testNestedModel: testNestedModel{
			Age: 11,
		},
	}
	age, ok := Property("age")(item).(int)
	if !(ok && age == 11) {
		t.Error("err")
	}
}
