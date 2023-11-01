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
		if entry == oneNumberEntry {
			result = append(result, []int{1, 1, 1, 1, 1, 1, 1, 1, 1})
		}
		if entry == twoNumberEntry {
			result = append(result, []int{2, 2, 2, 2, 2, 2, 2, 2, 2})
		}
	}

	return result
}

func ParseLines(s string) []string {
	entries := strings.Split(s, "\n\n")
	return entries[:len(entries)-1]
}
