package underscore

import (
	"testing"
)

func TestPropertyRV(t *testing.T) {
	item := TestModel{1, "one"}

	getAgeRV := PropertyRV("age")
	_, err := getAgeRV(item)
	if err == nil {
		t.Fatal("wrong")
	}

	getNameRV := PropertyRV("name")
	nameRV, err := getNameRV(item)
	if !(err == nil && nameRV.String() == item.Name) {
		t.Error("wrong")
	}
}

func TestProperty(t *testing.T) {
	item := TestModel{1, "one"}

	getAge := Property("age")
	_, err := getAge(item)
	if err == nil {
		t.Fatal("wrong")
	}

	getName := Property("name")
	name, err := getName(item)
	if !(err == nil && name.(string) == item.Name) {
		t.Error("wrong")
	}
}

func TestProperty_Ptr(t *testing.T) {
	item := &TestModel{1, "ptr"}

	nameGetter := Property("name")
	name, err := nameGetter(item)
	if err != nil || name != item.Name {
		t.Error(name, err)
	}
}
