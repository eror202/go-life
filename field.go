package life

import (
	"fmt"
	"strings"
)

const (
	deadCellRepresentation  = '.'
	aliveCellRepresentation = 'O'
)

type Field [][]bool

func NewField(width int, height int) Field {
	field := make(Field, height)
	for index := range field {
		field[index] = make([]bool, width)
	}

	return field
}

func ParseField(text string) (Field, error) {
	var field Field
	lines := strings.Split(text, "\n")
	width := -1
	for lineIndex, line := range lines {
		if line == "" || line[0] == '!' {
			continue
		}

		if width == -1 {
			width = len(line)
		} else if len(line) != width {
			return nil, fmt.Errorf("inconsistent length of line %d", lineIndex+1)
		}

		var row []bool
		for characterIndex, character := range line {
			switch character {
			case deadCellRepresentation:
				row = append(row, false)
			case aliveCellRepresentation:
				row = append(row, true)
			default:
				return nil, fmt.Errorf(
					"unexpected character %q at line %d and column %d",
					character,
					lineIndex+1,
					characterIndex+1,
				)
			}
		}

		field = append(field, row)
	}

	return field, nil
}

func (field Field) Width() int {
	return len(field[0])
}

func (field Field) Height() int {
	return len(field)
}

func (field Field) Cell(column int, row int) bool {
	column = wrapAroundModulus(column, field.Width())
	row = wrapAroundModulus(row, field.Height())
	return field[row][column]
}

func (field Field) SetCell(column int, row int, cell bool) {
	field[row][column] = cell
}

func (field Field) NeighborCount(column int, row int) int {
	var count int
	for rowOffset := -1; rowOffset <= 1; rowOffset++ {
		for columnOffset := -1; columnOffset <= 1; columnOffset++ {
			if rowOffset == 0 && columnOffset == 0 {
				continue
			}

			if field.Cell(column+columnOffset, row+rowOffset) {
				count++
			}
		}
	}

	return count
}

func (field Field) NextCell(column int, row int) bool {
	cell := field.Cell(column, row)
	neighborCount := field.NeighborCount(column, row)
	willBeBorn := !cell && neighborCount == 3
	willSurvive := cell && (neighborCount == 2 || neighborCount == 3)
	return willBeBorn || willSurvive
}

func (field Field) NextField() Field {
	nextField := NewField(field.Width(), field.Height())
	for row := 0; row < field.Height(); row++ {
		for column := 0; column < field.Width(); column++ {
			nextCell := field.NextCell(column, row)
			nextField.SetCell(column, row, nextCell)
		}
	}

	return nextField
}

func (field Field) String() string {
	result := ""
	for row := 0; row < field.Height(); row++ {
		for column := 0; column < field.Width(); column++ {
			if field.Cell(column, row) {
				result += string(aliveCellRepresentation)
			} else {
				result += string(deadCellRepresentation)
			}
		}

		result += "\n"
	}

	return result
}

func wrapAroundModulus(value int, modulus int) int {
	return (value + modulus) % modulus
}
