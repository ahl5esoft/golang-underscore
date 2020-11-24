package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Size(t *testing.T) {
	src := []string{"a", "b", "c"}
	res := Chain(src).Size()
	assert.Equal(t, res, 3)
}
