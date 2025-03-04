package model_test

import (
	"testing"

	"backend/second_left_right_equal/model"

	"github.com/stretchr/testify/assert"
)

func TestPossibleValue_AddPossibleValueAt(t *testing.T) {
	t.Run("test add possible value at 9", func(t *testing.T) {
		number := 9
		possible_value := model.NewPossibleNumberOf(number)

		assert.Equal(t, possible_value.Number, number)
		assert.Equal(t, possible_value.SmallestSum, 0)
		assert.Equal(t, len(possible_value.ValueOfSmallestSum), 0)
		assert.Equal(t, len(possible_value.PossibleValues), 10)
	})
}
