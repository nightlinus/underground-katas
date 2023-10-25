package bank_ocr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"underground-katas/bank_ocr"
)

func Test_read_empty_account_numbers_list(t *testing.T) {
	result := bank_ocr.ParseNumbers(
		`
	

	
	
	`)

	assert.Len(t, result, 0)
}
