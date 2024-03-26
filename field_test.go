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

func TestField_Cell(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, false, true},
	}
	cell := field.Cell(2, 1)

	assert.True(test, cell)
}

func TestField_Cell_withCoordinatesBeyondMinimum(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, true, false},
	}
	cell := field.Cell(-2, -1)

	assert.True(test, cell)
}

func TestField_Cell_withCoordinatesBeyondMaximum(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, true, false},
	}
	cell := field.Cell(4, 3)

	assert.True(test, cell)
}

func TestField_SetCell(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, false, false},
	}
	field.SetCell(2, 1, true)

	expectedField := Field{
		{false, false, false},
		{false, false, true},
	}
	assert.Equal(test, expectedField, field)
}
