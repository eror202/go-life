package life

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestField_Width(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, false, false},
	}
	result := field.Width()

	assert.Equal(test, 3, result)
}

func TestField_Height(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, false, false},
	}
	result := field.Height()

	assert.Equal(test, 2, result)
}
