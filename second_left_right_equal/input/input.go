package input

type Input struct {
	Format      string
	TotalDigits int
	Indice      []*InputIndex
}

type InputIndex struct {
	Value       string
	FirstIndex  int
	SecondIndex int
}

func NewInput(format string) *Input {
	input := Input{
		Format:      format,
		TotalDigits: len(format) + 1,
	}

	for i := 0; i < len(format); i++ {
		input.Indice = append(input.Indice, &InputIndex{
			Value:       string(format[i]),
			FirstIndex:  i,
			SecondIndex: i + 1,
		})
	}
	return &input
}
