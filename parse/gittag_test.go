package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	actual := Parse("v1.0.0")
	assert.Equal(t, 1, actual.Major)
	assert.Equal(t, 0, actual.Minor)
	assert.Equal(t, 0, actual.Patch)
}
