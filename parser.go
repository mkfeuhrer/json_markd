// json_markd helps you create json from markdown lists.
//
// Useful for creating API docs, personal todo, which can be used as json anywhere.
package json_markd

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Allowed markdown element types
type DataType int

const (
	Object  DataType = 0 // use object in markdown
	String  DataType = 1 // use string in markdown
	Integer DataType = 2 // use integer in markdown
	Double  DataType = 3 // use double in markdown
	Array   DataType = 4 // use array in markdown
	Invalid DataType = 5
)

var tabSpacesValue = 2

// SetTabSpaceValue is to set spaces used for tab in markdown. Eg. 2,4
func SetTabSpaceValue(val int) {
	tabSpacesValue = val
}

// ParseMarkdown accepts a filepath to a markdown file and return JSON string for same
func ParseMarkdown(filepath string) (string, error) {
	SetupLogger()
	markdownData, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
		return "", err
	}

	markdownDataList := strings.Split(strings.TrimSuffix(string(markdownData), "\n"), "\n")
	var lineDataList []string
	for _, line := range markdownDataList {
		list := strings.Split(line, "\t")
		if len(list) != 0 {
			lineDataList = append(lineDataList, list[0])
		}
	}
	markdownBlockList, err := createMarkdownBlockList(lineDataList)
	if err != nil {
		return "", err
	}
	result := generateMarkdownString(markdownBlockList)
	return result, nil
}

func generateMarkdownString(markdownBlockList []markdownBlock) string {
	result := "{\n"
	s := itemStack{}
	s.new()
	for ind, markdownBlock := range markdownBlockList {
		// if at last element
		if (ind + 1) == len(markdownBlockList) {
			topIndex := 0
			if s.size() > 0 {
				topIndex = (*s.top()).(int)
			}
			if markdownBlockList[topIndex].Value == Array {
				result += strings.Repeat("  ", markdownBlock.TabCount+1) + markdownBlock.GetPrefixForDatatypeWhenParentIsArray() + markdownBlock.GetSuffixForDatatype() + "\n"
			} else {
				result += strings.Repeat("  ", markdownBlock.TabCount+1) + markdownBlock.GetPrefixForDatatype() + markdownBlock.GetSuffixForDatatype() + "\n"
			}
			for s.size() > 0 {
				item, _ := s.pop()
				index := (*item).(int)
				result += strings.Repeat("  ", markdownBlockList[index].TabCount+1) + markdownBlockList[index].GetSuffixForDatatype() + "\n"
			}
			break
		}
		// if nested block present
		if markdownBlock.TabCount < markdownBlockList[ind+1].TabCount {
			// append start string to result
			topIndex := 0
			if s.size() > 0 {
				topIndex = (*s.top()).(int)
			}
			if markdownBlockList[topIndex].Value == Array {
				result += strings.Repeat("  ", markdownBlock.TabCount+1) + markdownBlock.GetPrefixForDatatypeWhenParentIsArray()
			} else {
				result += strings.Repeat("  ", markdownBlock.TabCount+1) + markdownBlock.GetPrefixForDatatype()
			}
			s.push(ind)
		} else {
			topIndex := 0
			if s.size() > 0 {
				topIndex = (*s.top()).(int)
			}
			if markdownBlock.TabCount == markdownBlockList[ind+1].TabCount {
				if markdownBlockList[topIndex].Value == Array {
					result += strings.Repeat("  ", markdownBlock.TabCount+1) + markdownBlock.GetPrefixForDatatypeWhenParentIsArray() + markdownBlock.GetSuffixForDatatype() + ",\n"
				} else {
					result += strings.Repeat("  ", markdownBlock.TabCount+1) + markdownBlock.GetPrefixForDatatype() + markdownBlock.GetSuffixForDatatype() + ",\n"
				}
			} else if markdownBlock.TabCount > markdownBlockList[ind+1].TabCount {
				if markdownBlockList[topIndex].Value == Array {
					result += strings.Repeat("  ", markdownBlock.TabCount+1) + markdownBlock.GetPrefixForDatatypeWhenParentIsArray() + markdownBlock.GetSuffixForDatatype() + "\n"
				} else {
					result += strings.Repeat("  ", markdownBlock.TabCount+1) + markdownBlock.GetPrefixForDatatype() + markdownBlock.GetSuffixForDatatype() + "\n"
				}
			}
			for s.size() > 0 {
				index := (*s.top()).(int)
				if markdownBlockList[index].TabCount == markdownBlockList[ind+1].TabCount {
					result += strings.Repeat("  ", markdownBlockList[index].TabCount+1) + markdownBlockList[index].GetSuffixForDatatype() + ",\n"
					s.pop()
				} else if markdownBlockList[index].TabCount > markdownBlockList[ind+1].TabCount {
					result += strings.Repeat("  ", markdownBlockList[index].TabCount+1) + markdownBlockList[index].GetSuffixForDatatype() + "\n"
					s.pop()
				} else {
					break
				}
			}
		}
	}
	result += "}"
	return result
}

func createMarkdownBlockList(lineDataList []string) ([]markdownBlock, error) {
	var markdownBlockList []markdownBlock
	for _, line := range lineDataList {
		tabCount, lineWithoutTabs := removeTabsFromLines(line, tabSpacesValue)
		key, val, err := parseLine(lineWithoutTabs)
		if err != nil {
			return []markdownBlock{}, err
		}
		dataType := getDatatypeFromVal(val)
		markdownBlock := newMarkdownBlock(tabCount, key, dataType)
		markdownBlockList = append(markdownBlockList, markdownBlock)
	}
	return markdownBlockList, nil
}
