package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsMatch(t *testing.T) {
	assert.False(
		t,
		IsMatch(nil, nil),
	)

	m := testModel{ID: 1, Name: "one"}
	assert.False(
		t,
		IsMatch(m, nil),
	)

	assert.False(
		t,
		IsMatch(m, map[string]interface{}{
			"id":   m.ID,
			"name": "a",
		}),
	)

	assert.True(
		t,
		IsMatch(m, map[string]interface{}{
			"id":   m.ID,
			"name": m.Name,
		}),
	)
}
