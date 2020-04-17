package json_markd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMarkdownBlockList(t *testing.T) {
	SetTabSpaceValue(2)
	SetupLogger()
	t.Run("when correct data list is passed", func(t *testing.T) {
		lineDataList := []string{
			" - data: object",
		}
		t.Run("it should return correct response", func(t *testing.T) {
			expectedResponse := []MarkdownBlock{MarkdownBlock{TabCount: 0, Key: "\"data\"", Value: 0}}
			markdownBlock, _ := createMarkdownBlockList(lineDataList)
			assert.Equal(t, expectedResponse, markdownBlock)
		})
	})

	t.Run("when incorrect data list is passed", func(t *testing.T) {
		lineDataList := []string{
			"data object",
		}
		SetTabSpaceValue(2)
		t.Run("it should return error", func(t *testing.T) {
			expectedResponse := errors.New(".errors.invalid_markdown_list_format")
			_, err := createMarkdownBlockList(lineDataList)
			assert.Equal(t, expectedResponse, err)
		})
	})
}

func TestParseMarkdown(t *testing.T) {
	SetTabSpaceValue(2)
	SetupLogger()
	t.Run("when correct file is passed", func(t *testing.T) {
		t.Run("it should return correct response", func(t *testing.T) {
			expectedResponse := "{\n  \"data\" : {\n    \"name\" : \"random string\",\n    \"age\" : 0,\n    \"income\" : 0.0,\n    \"vehicles\" : [\n      {\n        \"name\" : \"random string\",\n        \"price\" : 0.0\n      }\n    ],\n    \"apps\" : [\n      [\n        \"random string\",\n        \"random string\"\n      ],\n      [\n        \"random string\",\n        \"random string\"\n      ]\n    ]\n  },\n  \"errors\" : {\n    \"type\" : \"random string\"\n  }\n}"
			result, _ := ParseMarkdown("./test_data/sample.md")
			assert.Equal(t, expectedResponse, result)
		})
	})

	t.Run("when incorrect file is passed", func(t *testing.T) {
		t.Run("it should return correct response", func(t *testing.T) {
			expectedResponse := "open ./test_data/sample_1.md: no such file or directory"
			_, err := ParseMarkdown("./test_data/sample_1.md")
			assert.Equal(t, expectedResponse, err.Error())
		})
	})

	t.Run("when file data is invalid", func(t *testing.T) {
		t.Run("it should return correct response", func(t *testing.T) {
			expectedResponse := ".errors.invalid_markdown_list_format"
			_, err := ParseMarkdown("./test_data/sample_invalid.md")
			assert.Equal(t, expectedResponse, err.Error())
		})
	})
}
