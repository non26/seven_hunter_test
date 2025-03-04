package model

type PossibleNumber struct {
	Number             int
	SmallestSum        int
	ValueOfSmallestSum []int
	PossibleValues     []*PossibleValue
}

type PossibleValue struct {
	Value []int
	Sum   int
}

func NewPossibleValue(value []int, sum int) *PossibleValue {
	return &PossibleValue{
		Value: value,
		Sum:   sum,
	}
}

func (p *PossibleValue) SumOfValue(another int) {
	p.Sum = p.Sum + another
}

func (p *PossibleValue) AddValue(another int) {
	p.Value = append(p.Value, another)
}

func NewPossibleNumberOf(number int) *PossibleNumber {
	p := &PossibleNumber{
		Number:         number,
		PossibleValues: make([]*PossibleValue, 0),
	}
	p.InitializePossibleValues()
	return p
}

//	func (p *PossibleNumber) InitializePossibleValues() {
//		numberOfArray := 10
//		for i := 0; i < numberOfArray; i++ {
//			value := make([]int, 1)
//			value[0] = p.Number
//			p.PossibleValues = append(p.PossibleValues, &PossibleValue{
//				Value: value,
//				Sum:   p.Number,
//			})
//		}
//	}

func (p *PossibleNumber) InitializePossibleValues() {
	value := make([]int, 1)
	value[0] = p.Number
	p.PossibleValues = append(p.PossibleValues, &PossibleValue{
		Value: value,
		Sum:   p.Number,
	})
}

func (p *PossibleNumber) TrimPossibleValues(largestIndex int) {
	if largestIndex == 0 {
		p.PossibleValues = []*PossibleValue{}
	} else {
		p.PossibleValues = p.PossibleValues[0:largestIndex]
	}
}

func (p *PossibleNumber) GetSmallestSum() (int, []int) {
	smallestSum := p.PossibleValues[0].Sum
	smallestSumValue := p.PossibleValues[0].Value
	for _, possibleValue := range p.PossibleValues {
		if possibleValue.Sum < smallestSum {
			smallestSum = possibleValue.Sum
			smallestSumValue = possibleValue.Value
		}
	}
	return smallestSum, smallestSumValue
}
