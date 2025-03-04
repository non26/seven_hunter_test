package solution_test

import (
	"backend/second_left_right_equal/input"
	"backend/second_left_right_equal/solution"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Run("LLRR=", func(t *testing.T) {
		input := input.NewInput("LLRR=")
		actual := solution.Solution(input)
		expected := "210122"
		assert.Equal(t, expected, actual)
	})
	t.Run("==RLL", func(t *testing.T) {
		input := input.NewInput("==RLL")
		actual := solution.Solution(input)
		expected := "000210"
		assert.Equal(t, expected, actual)
	})

	t.Run("=LLRR", func(t *testing.T) {
		input := input.NewInput("=LLRR")
		actual := solution.Solution(input)
		expected := "221012"
		assert.Equal(t, expected, actual)
	})

	t.Run("RRL=R", func(t *testing.T) {
		input := input.NewInput("RRL=R")
		actual := solution.Solution(input)
		expected := "012001"
		assert.Equal(t, expected, actual)
	})

}
