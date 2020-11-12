package underscore

import (
	"testing"
)

func Benchmark_Property(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PropertyRV("id")(testModel{ID: 1, Name: "one"})
	}
}

func Benchmark_PropertyRV(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PropertyRV("id")(testModel{ID: 1, Name: "one"})
	}
}

func Benchmark_PropertyRV_InvalidName(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PropertyRV("$$")(testModel{ID: 1, Name: "one"})
	}
}

func Benchmark_PropertyRV_Nested(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PropertyRV("id")(testNestedModel{
			testModel: testModel{
				ID: 11,
			},
		})
	}
}

func Benchmark_PropertyRV_Ptr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PropertyRV("id")(&testModel{ID: 1, Name: "ptr"})
	}
}

func Test_Property(t *testing.T) {
	item := testModel{ID: 1, Name: "one"}
	getName := Property("name")
	name := getName(item)
	if name.(string) != item.Name {
		t.Error("wrong")
	}
}

func Test_Property_Nested(t *testing.T) {
	item := testNestedModel{
		testModel: testModel{
			ID: 11,
		},
	}
	id, ok := Property("id")(item).(int)
	if !(ok && id == 11) {
		t.Error("err")
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

func Test_PropertyRV(t *testing.T) {
	item := testModel{ID: 1, Name: "one"}

	rv := PropertyRV("$$")(item)
	if rv != nilValue {
		t.Fatal("wrong")
	}

	getNameRV := PropertyRV("name")
	nameRV := getNameRV(item)
	if nameRV.String() != item.Name {
		t.Error("wrong")
	}
}
