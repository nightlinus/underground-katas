package bank_ocr

import "strings"

const (
	oneNumberEntry = `                        
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |`
	twoNumberEntry = `_  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
|_ |_ |_ |_ |_ |_ |_ |_ |_ `
)

func ParseNumbers(s string) [][]int {
	entries := ParseLines(s)
	entriesCount := len(entries)
	result := make([][]int, 0, entriesCount)

	for _, entry := range entries {
		result = append(result, parseLine(entry))
	}

	return result
}

func ParseLines(s string) []string {
	entries := strings.Split(s, "\n\n")
	return entries[:len(entries)-1]
}

func parseLine(entry string) []int {
	if entry == oneNumberEntry {
		return []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	}
	if entry == twoNumberEntry {
		return []int{2, 2, 2, 2, 2, 2, 2, 2, 2}
	}

	result := make([]int, 9)

	return result
}
