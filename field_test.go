package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_Field(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Field("id")(testModel{ID: 1, Name: "one"})
	}
}

func Benchmark_FieldValue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FieldValue("id")(testModel{ID: 1, Name: "one"})
	}
}

func Benchmark_FieldValue_InvalidName(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FieldValue("$$")(testModel{ID: 1, Name: "one"})
	}
}

func Benchmark_FieldValue_Nested(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FieldValue("id")(testNestedModel{
			testModel: testModel{
				ID: 11,
			},
		})
	}
}

func Benchmark_FieldValue_Ptr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FieldValue("id")(&testModel{ID: 1, Name: "ptr"})
	}
}

func Test_Field(t *testing.T) {
	item := testModel{ID: 1, Name: "one"}
	getName := Field("name")
	name := getName(item)
	assert.Equal(
		t,
		name,
		item.Name,
	)
}

func Test_Field_Nested(t *testing.T) {
	item := testNestedModel{
		testModel: testModel{
			ID: 11,
		},
	}
	id, ok := Field("id")(item).(int)
	assert.True(t, ok)
	assert.Equal(t, id, 11)
}

func Test_Field_Ptr(t *testing.T) {
	item := &testModel{ID: 1, Name: "ptr"}
	nameGetter := Field("name")
	name := nameGetter(item)
	assert.Equal(t, name, item.Name)
}

func Test_FieldValue(t *testing.T) {
	item := testModel{ID: 1, Name: "one"}

	value := FieldValue("$$")(item)
	assert.Equal(t, value, nilValue)

	getNameValue := FieldValue("name")
	nameValue := getNameValue(item)
	assert.Equal(
		t,
		nameValue.String(),
		item.Name,
	)
}
