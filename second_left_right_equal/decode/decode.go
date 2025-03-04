package decode

import (
	"backend/second_left_right_equal/input"
	"backend/second_left_right_equal/model"
	"math"
)

func Decode(initial_possible_number *model.PossibleNumber, format string, input *input.Input) {
	switch format {
	case "L":
		DecodeLeft(initial_possible_number, input)
	case "R":
		DecodeRight(initial_possible_number, input)
	case "=":
		DecodeEqual(initial_possible_number, input)
	}
}

func DecodeLeft(possible_number *model.PossibleNumber, input *input.Input) {
	resultPossibleValues := []*model.PossibleValue{}
	smallestSum := math.MaxInt
	smallestSumValue := []int{}
	for _, possibleValue := range possible_number.PossibleValues {
		targetValue := possibleValue.Value[len(possibleValue.Value)-1]
		for i := targetValue - 1; 0 <= i; i-- {
			value := make([]int, len(possibleValue.Value))
			copy(value, possibleValue.Value)
			newPossibleValue := model.NewPossibleValue(value, possibleValue.Sum)
			newPossibleValue.AddValue(i)
			newPossibleValue.SumOfValue(i)
			resultPossibleValues = append(resultPossibleValues, newPossibleValue)
			if len(newPossibleValue.Value) == input.TotalDigits && newPossibleValue.Sum < smallestSum {
				smallestSum = newPossibleValue.Sum
				smallestSumValue = newPossibleValue.Value
			}
		}
	}
	possible_number.PossibleValues = resultPossibleValues
	possible_number.SmallestSum = smallestSum
	possible_number.ValueOfSmallestSum = smallestSumValue
}

func DecodeRight(possible_number *model.PossibleNumber, input *input.Input) {
	// if possible_number.PossibleValues[0].Value[0] == 0 {
	// 	possible_number.TrimPossibleValues(0)
	// 	return
	// }
	smallestSum := math.MaxInt
	smallestSumValue := []int{}
	resultPossibleValues := []*model.PossibleValue{}
	for _, possibleValue := range possible_number.PossibleValues {
		targetValue := possibleValue.Value[len(possibleValue.Value)-1]
		for i := targetValue + 1; i <= 9; i++ {
			value := make([]int, len(possibleValue.Value))
			copy(value, possibleValue.Value)
			newPossibleValue := model.NewPossibleValue(value, possibleValue.Sum)
			newPossibleValue.AddValue(i)
			newPossibleValue.SumOfValue(i)
			resultPossibleValues = append(resultPossibleValues, newPossibleValue)
			if len(newPossibleValue.Value) == input.TotalDigits && newPossibleValue.Sum < smallestSum {
				smallestSum = newPossibleValue.Sum
				smallestSumValue = newPossibleValue.Value
			}
		}
	}
	possible_number.PossibleValues = resultPossibleValues
	possible_number.SmallestSum = smallestSum
	possible_number.ValueOfSmallestSum = smallestSumValue
}

func DecodeEqual(possible_number *model.PossibleNumber, input *input.Input) {
	// if possible_number.PossibleValues[0].Value[0] == 0 {
	// 	possible_number.TrimPossibleValues(0)
	// 	return
	// }
	smallestSum := math.MaxInt
	smallestSumValue := []int{}
	resultPossibleValues := []*model.PossibleValue{}
	for _, possibleValue := range possible_number.PossibleValues {
		targetValue := possibleValue.Value[len(possibleValue.Value)-1]
		value := make([]int, len(possibleValue.Value))
		copy(value, possibleValue.Value)
		newPossibleValue := model.NewPossibleValue(value, possibleValue.Sum)
		newPossibleValue.AddValue(targetValue)
		newPossibleValue.SumOfValue(targetValue)
		resultPossibleValues = append(resultPossibleValues, newPossibleValue)
		if len(newPossibleValue.Value) == input.TotalDigits && newPossibleValue.Sum < smallestSum {
			smallestSum = newPossibleValue.Sum
			smallestSumValue = newPossibleValue.Value
		}
	}
	possible_number.PossibleValues = resultPossibleValues
	possible_number.SmallestSum = smallestSum
	possible_number.ValueOfSmallestSum = smallestSumValue
}
