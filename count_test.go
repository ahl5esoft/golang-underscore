package underscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Count(t *testing.T) {
	src := []string{"a", "b", "c"}
	res := Chain(src).Count()
	assert.Len(t, src, res)
}
