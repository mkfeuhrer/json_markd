package utils

import (
	"errors"
	"json-markd/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveTabsFromLines(t *testing.T) {
	expectedResult := "hello"
	expectedTabCount := 2
	actualTabCount, actualResult := RemoveTabsFromLines("    hello", 2)
	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedTabCount, actualTabCount)
}

func TestTrimString(t *testing.T) {
	expectedResult := "hello"
	actualResult := TrimString(" hello  ", " ")
	assert.Equal(t, expectedResult, actualResult)
}

func TestParseLine(t *testing.T) {
	logger.SetupLogger()
	t.Run("when correct input data is passed", func(t *testing.T) {
		t.Run("it should return correct result", func(t *testing.T) {
			expectedKey := "\"data\""
			expectedVal := "object"
			actualKey, actualVal, err := ParseLine("- data: object")
			assert.Equal(t, expectedKey, actualKey)
			assert.Equal(t, expectedVal, actualVal)
			assert.NoError(t, err)
		})
	})

	t.Run("when incorrect input data is passed", func(t *testing.T) {
		t.Run("it should return error", func(t *testing.T) {
			expectedError := errors.New(".errors.invalid_markdown_list_format")
			_, _, err := ParseLine("- data object")
			assert.Equal(t, expectedError, err)
		})
	})
}
