package bank_ocr_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"underground-katas/bank_ocr"
)

func genNLines(n int) string {
	return strings.Repeat(`   
  |
  |

`, n)
}

func Test_read_empty_account_numbers_list(t *testing.T) {
	result := bank_ocr.ParseNumbers(``)

	assert.Len(t, result, 0)
}

func Test_recognize_first_line(t *testing.T) {
	result := bank_ocr.ParseNumbers(genNLines(1))
	assert.Len(t, result, 1)
}

func Test_recognize_two_line(t *testing.T) {
	result := bank_ocr.ParseNumbers(genNLines(2))
	assert.Len(t, result, 2)
}

func Test_recognize_500_line(t *testing.T) {
	result := bank_ocr.ParseNumbers(genNLines(500))
	assert.Len(t, result, 500)
}
