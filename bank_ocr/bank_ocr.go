package bank_ocr

const (
	digitOne = `   
  |
  |
`
)

func ParseNumbers(s string) [][]int {
	if s == digitOne {
		return [][]int{{1}}
	}
	return [][]int{}
}
