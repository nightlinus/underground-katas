package bank_ocr

import "strings"

const (
	oneNumber = `   
  |
  |

`
)

func ParseNumbers(s string) [][]int {
	i := strings.Count(s, oneNumber)
	result := make([][]int, i)

	return result
}
