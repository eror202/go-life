package life

type Field [][]bool

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

func wrapAroundModulus(value int, modulus int) int {
	return (value + modulus) % modulus
}
