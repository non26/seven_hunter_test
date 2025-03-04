package main

import (
	"backend/second_left_right_equal/input"
	"backend/second_left_right_equal/solution"
)

func main() {
	user_input := "LLRR="
	input := input.NewInput(user_input)
	solution.Solution(input)
}
