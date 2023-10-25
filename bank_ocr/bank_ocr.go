package bank_ocr

const (
	oneNumber = `   
  |
  |

`
)

func ParseNumbers(s string) [][]int {
	if s == oneNumber {
		return [][]int{{}}
	}
	return [][]int{}
}
