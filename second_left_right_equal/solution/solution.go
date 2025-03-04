package solution

import (
	"backend/second_left_right_equal/decode"
	"backend/second_left_right_equal/input"
	"backend/second_left_right_equal/model"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solution(input *input.Input) string {
	smallestSum := math.MaxInt
	smallestSumValue := []int{}
	for i := 0; i <= 9; i++ {
		possibleNumber := model.NewPossibleNumberOf(i)
		for _, index := range input.Indice {
			decode.Decode(possibleNumber, index.Value, input)
			if len(possibleNumber.PossibleValues) == 0 {
				break
			}
		}
		if possibleNumber.SmallestSum < smallestSum {
			smallestSum = possibleNumber.SmallestSum
			smallestSumValue = possibleNumber.ValueOfSmallestSum
		}
	}
	strArr := make([]string, len(smallestSumValue))
	for i, num := range smallestSumValue {
		strArr[i] = strconv.Itoa(num) // Convert int to string
	}
	str := strings.Join(strArr, "")
	fmt.Printf("smallestSum: %d, smallestSumValue: %v", smallestSum, str)
	return str
}
