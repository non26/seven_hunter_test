package findvalue_test

import (
	findvalue "backend/first_most_value/find_value"
	readinput "backend/first_most_value/read_input"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindValue(t *testing.T) {
	t.Run("read json: test_case_3.json", func(t *testing.T) {
		jsonFile := "test_case/test_case_3.json"
		input, err := readinput.ReadInput(jsonFile)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		expected := 237
		actual := findvalue.FindValue(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("read json: test_case_4.json", func(t *testing.T) {
		jsonFile := "test_case/test_case_4.json"
		input, err := readinput.ReadInput(jsonFile)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		expected := 7273
		actual := findvalue.FindValue(input)
		assert.Equal(t, expected, actual)
	})
}
