package underscore

import (
	"testing"
)

func TestPropertyRV(t *testing.T) {
	item := TestModel{ 1, "one" }
	
	getAgeRV := PropertyRV("age")
	_, err := getAgeRV(item)
	if err == nil {
		t.Error("wrong")
		return
	}

	getNameRV := PropertyRV("name")
	nameRV, err := getNameRV(item)
	if !(err == nil && nameRV.String() == item.Name) {
		t.Error("wrong")
	}
}

func TestProperty(t *testing.T) {
	item := TestModel{ 1, "one" }
	
	getAge := Property("age")
	_, err := getAge(item)
	if err == nil {
		t.Error("wrong")
		return
	}

	getName := Property("name")
	name, err := getName(item)
	if !(err == nil && name.(string) == item.Name) {
		t.Error("wrong")
	}
}