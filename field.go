package life

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
