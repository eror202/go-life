package life

type Field [][]bool

func (field Field) Width() int {
	return len(field[0])
}

func (field Field) Height() int {
	return len(field)
}
