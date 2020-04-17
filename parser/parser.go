package parser

import (
	"fmt"
	"io/ioutil"
	"json-markd/logger"
	"json-markd/stack"
	"json-markd/utils"
	"strings"
)

type DataType int

const (
	Object  DataType = 0
	String  DataType = 1
	Integer DataType = 2
	Double  DataType = 3
	Array   DataType = 4
	Invalid DataType = 4
)

var tabSpacesValue = 2

func SetTabSpaceValue(val int) {
	tabSpacesValue = val
}

func ParseMarkdown(filepath string) (string, error) {
	logger.SetupLogger()
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

func generateMarkdownString(markdownBlockList []MarkdownBlock) string {
	result := "{\n"
	s := stack.ItemStack{}
	s.New()
	for ind, markdownBlock := range markdownBlockList {
		// if at last element
		if (ind + 1) == len(markdownBlockList) {
			topIndex := 0
			if s.Size() > 0 {
				topIndex = (*s.Top()).(int)
			}
			if markdownBlockList[topIndex].Value == Array {
				result += strings.Repeat("  ", markdownBlock.TabCount+1) + markdownBlock.GetPrefixForDatatypeWhenParentIsArray() + markdownBlock.GetSuffixForDatatype() + "\n"
			} else {
				result += strings.Repeat("  ", markdownBlock.TabCount+1) + markdownBlock.GetPrefixForDatatype() + markdownBlock.GetSuffixForDatatype() + "\n"
			}
			for s.Size() > 0 {
				item, _ := s.Pop()
				index := (*item).(int)
				result += strings.Repeat("  ", markdownBlockList[index].TabCount+1) + markdownBlockList[index].GetSuffixForDatatype() + "\n"
			}
			break
		}
		// if nested block present
		if markdownBlock.TabCount < markdownBlockList[ind+1].TabCount {
			// append start string to result
			topIndex := 0
			if s.Size() > 0 {
				topIndex = (*s.Top()).(int)
			}
			if markdownBlockList[topIndex].Value == Array {
				result += strings.Repeat("  ", markdownBlock.TabCount+1) + markdownBlock.GetPrefixForDatatypeWhenParentIsArray()
			} else {
				result += strings.Repeat("  ", markdownBlock.TabCount+1) + markdownBlock.GetPrefixForDatatype()
			}
			s.Push(ind)
		} else {
			topIndex := 0
			if s.Size() > 0 {
				topIndex = (*s.Top()).(int)
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
			for s.Size() > 0 {
				index := (*s.Top()).(int)
				if markdownBlockList[index].TabCount == markdownBlockList[ind+1].TabCount {
					result += strings.Repeat("  ", markdownBlockList[index].TabCount+1) + markdownBlockList[index].GetSuffixForDatatype() + ",\n"
					s.Pop()
				} else if markdownBlockList[index].TabCount > markdownBlockList[ind+1].TabCount {
					result += strings.Repeat("  ", markdownBlockList[index].TabCount+1) + markdownBlockList[index].GetSuffixForDatatype() + "\n"
					s.Pop()
				} else {
					break
				}
			}
		}
	}
	result += "}"
	return result
}

func createMarkdownBlockList(lineDataList []string) ([]MarkdownBlock, error) {
	var markdownBlockList []MarkdownBlock
	for _, line := range lineDataList {
		tabCount, lineWithoutTabs := utils.RemoveTabsFromLines(line, tabSpacesValue)
		key, val, err := utils.ParseLine(lineWithoutTabs)
		if err != nil {
			return []MarkdownBlock{}, err
		}
		dataType := getDatatypeFromVal(val)
		markdownBlock := NewMarkdownBlock(tabCount, key, dataType)
		markdownBlockList = append(markdownBlockList, markdownBlock)
	}
	return markdownBlockList, nil
}
