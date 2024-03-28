package life

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewField(test *testing.T) {
	field := NewField(3, 2)

	expectedField := Field{
		{false, false, false},
		{false, false, false},
	}
	assert.Equal(test, expectedField, field)
}

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

func TestField_NeighborCount(test *testing.T) {
	field := Field{
		{true, false, false},
		{false, true, true},
		{true, false, false},
	}
	count := field.NeighborCount(0, 1)

	assert.Equal(test, 4, count)
}

func TestField_NeighborCount_forAliveCell(test *testing.T) {
	field := Field{
		{true, false, false},
		{true, true, true},
		{true, false, false},
	}
	count := field.NeighborCount(0, 1)

	assert.Equal(test, 4, count)
}

func TestField_NextCell_willBeBorn(test *testing.T) {
	field := Field{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false /* ! */, false, true, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
	}
	cell := field.NextCell(1, 2)

	assert.True(test, cell)
}

func TestField_NextCell_willSurvive(test *testing.T) {
	field := Field{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, false, true /* ! */, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
	}
	cell := field.NextCell(3, 2)

	assert.True(test, cell)
}

func TestField_NextCell_willDie(test *testing.T) {
	field := Field{
		{false, false, false, false, false},
		{false, false, true /* ! */, false, false},
		{false, false, false, true, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
	}
	cell := field.NextCell(2, 1)

	assert.False(test, cell)
}

func TestField_NextField(test *testing.T) {
	field := Field{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, false, true, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
	}
	nextField := field.NextField()

	expectedNextField := Field{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, true, false, true, false},
		{false, false, true, true, false},
		{false, false, true, false, false},
	}
	assert.Equal(test, expectedNextField, nextField)
}

func TestField_String(test *testing.T) {
	field := Field{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}
	result := field.String()

	assert.Equal(test, ".O.\n..O\nOOO\n", result)
}
