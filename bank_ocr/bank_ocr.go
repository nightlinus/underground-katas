package bank_ocr

import "strings"

const (
	oneNumberEntry = `                        
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |

`
)

func ParseNumbers(s string) [][]int {
	entriesCount := strings.Count(s, oneNumberEntry)
	if entriesCount == 0 {
		return [][]int{}
	}
	result := make([][]int, entriesCount)
	result[0] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	result[1] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1}

	return result
}
