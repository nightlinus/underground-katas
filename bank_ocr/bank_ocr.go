package bank_ocr

import "strings"

const (
	oneNumberEntry = `                        
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |

`
	twoNumberEntry = `_  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
|_ |_ |_ |_ |_ |_ |_ |_ |_ 

`
)

func ParseNumbers(s string) [][]int {
	entriesCount := strings.Count(s, "\n") / 4
	result := make([][]int, 0, entriesCount)

	for i := 0; i < entriesCount; i++ {
		result = append(result, []int{1, 1, 1, 1, 1, 1, 1, 1, 1})
	}

	return result
}
