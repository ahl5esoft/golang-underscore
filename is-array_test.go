package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsArray(t *testing.T) {
	assert.True(
		t,
		IsArray([]int{}),
	)
	assert.False(
		t,
		IsArray(map[string]int{}),
	)
}
