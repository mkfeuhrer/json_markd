package json_markd

import (
	"errors"
	"strings"
)

func RemoveTabsFromLines(line string, tabSpacesValue int) (int, string) {
	count := 0
	for _, character := range line {
		if character == ' ' {
			count++
		} else {
			break
		}
	}
	count = count / tabSpacesValue
	lineWithoutTabs := strings.Trim(line, " ")
	return count, lineWithoutTabs
}

func TrimString(input string, sep string) string {
	return strings.Trim(input, sep)
}

func ParseLine(line string) (string, string, error) {
	line = strings.TrimLeft(line, "-")
	line = TrimString(line, " ")
	keyValueList := strings.Split(line, ":")
	if len(keyValueList) < 2 {
		Log.Error(".errors.invalid_markdown_list_format")
		return "", "", errors.New(".errors.invalid_markdown_list_format")
	}
	// return error if len(keyValueList) < 2
	return "\"" + TrimString(keyValueList[0], " ") + "\"", TrimString(keyValueList[1], " "), nil
}
