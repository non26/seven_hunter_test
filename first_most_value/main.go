package main

import (
	findvalue "backend/first_most_value/find_value"
	readinput "backend/first_most_value/read_input"
	"fmt"
)

func main() {
	input, err := readinput.ReadInput("test_case/test_case_4.json")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	result := findvalue.FindValue(input)
	fmt.Println("Result:", result)
}
