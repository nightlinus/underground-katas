package bank_ocr_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"underground-katas/bank_ocr"
)

func genNLines(n int) string {
	return strings.Repeat(`                           
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |

`, n)
}

func genNResults(n int) []bank_ocr.Account {
	result := make([]bank_ocr.Account, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, bank_ocr.MustAccount(1, 1, 1, 1, 1, 1, 1, 1, 1))
	}

	return result
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

func Test_recognize_with_diff_lines(t *testing.T) {
	lines := `_  _  _  _  _  _  _  _  _  
 _| _| _| _| _| _| _| _| _|
|_ |_ |_ |_ |_ |_ |_ |_ |_ 

`
	result := bank_ocr.ParseNumbers(lines)
	assert.Len(t, result, 1)
}

func Test_parse_line_with_mixed_symbols(t *testing.T) {
	lines := `   _  _  _  _  _  _  _  _  
  | _| _| _| _| _| _| _| _|
  | |_ |_ |_ |_ |_ |_ |_ |_ 

`
	result := bank_ocr.ParseNumbers(lines)
	assert.Len(t, result[0], 9)
}

func Test_recognize_full_line(t *testing.T) {
	expected := genNResults(1)
	result := bank_ocr.ParseNumbers(genNLines(1))
	assert.Equalf(t, expected, result,
		`ParseNumbers() want = %v, got = %v`, [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1}}, result)
}

func Test_recognize_full_lines(t *testing.T) {
	expected := genNResults(2)

	result := bank_ocr.ParseNumbers(genNLines(2))

	assert.Equalf(t, expected, result,
		`ParseNumbers() want = %v, got = %v`, expected, result)
}

func Test_recognize_diff_lines(t *testing.T) {
	expected := []bank_ocr.Account{bank_ocr.MustAccount(2, 2, 2, 2, 2, 2, 2, 2, 2)}

	result := bank_ocr.ParseNumbers(
		` _  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
|_ |_ |_ |_ |_ |_ |_ |_ |_ 

`)

	assert.Equalf(t, expected, result,
		`ParseNumbers() want = %v, got = %v`, expected, result)
}

func Test_parse_lines_with_1_and_2(t *testing.T) {
	result := bank_ocr.ParseNumbers(
		`    _  _  _  _  _  _  _  _ 
  | _| _| _| _| _| _| _| _|
  ||_ |_ |_ |_ |_ |_ |_ |_ 

`)

	expected := []bank_ocr.Account{bank_ocr.MustAccount(1, 2, 2, 2, 2, 2, 2, 2, 2)}
	assert.Equalf(t, expected, result,
		`ParseNumbers() want = %v, got = %v`, expected, result)
}

func Test_parse_lines_with_1_2_3(t *testing.T) {
	result := bank_ocr.ParseNumbers(
		`    _  _  _  _  _  _  _  _ 
  | _| _| _| _| _| _| _| _|
  ||_  _||_ |_ |_ |_ |_ |_ 

`)

	expected := []bank_ocr.Account{bank_ocr.MustAccount(1, 2, 3, 2, 2, 2, 2, 2, 2)}
	assert.Equalf(t, expected, result,
		`ParseNumbers() want = %v, got = %v`, expected, result)
}

func Test_parse_lines_with_all_numbers(t *testing.T) {
	result := bank_ocr.ParseNumbers(
		`    _  _     _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _|

`)

	expected := []bank_ocr.Account{bank_ocr.MustAccount(1, 2, 3, 4, 5, 6, 7, 8, 9)}
	assert.Equalf(t, expected, result,
		`ParseNumbers() want = %v, got = %v`, expected, result)
}

func Test_calculate_checksum_for_non_zero_first_position(t *testing.T) {
	checkSum := bank_ocr.CalculateCheckSum(bank_ocr.MustAccount(0, 0, 0, 0, 0, 0, 0, 0, 1))

	assert.Equal(t, 1, checkSum)
}

func Test_calculate_d2_coefficient(t *testing.T) {
	checkSum := bank_ocr.CalculateCheckSum(bank_ocr.MustAccount(0, 0, 0, 0, 0, 0, 0, 1, 0))

	assert.Equal(t, 2, checkSum)
}

func Test_calculate_d3_coefficient(t *testing.T) {
	checkSum := bank_ocr.CalculateCheckSum(bank_ocr.MustAccount(0, 0, 0, 0, 0, 0, 1, 0, 0))

	assert.Equal(t, 3, checkSum)
}

func Test_calculate_d9_coefficient(t *testing.T) {
	checkSum := bank_ocr.CalculateCheckSum(bank_ocr.MustAccount(1, 0, 0, 0, 0, 0, 0, 0, 0))

	assert.Equal(t, 9, checkSum)
}

func Test_calculate_d9_coefficient_value_2(t *testing.T) {
	checkSum := bank_ocr.CalculateCheckSum(bank_ocr.MustAccount(2, 0, 0, 0, 0, 0, 0, 0, 0))

	assert.Equal(t, 18, checkSum)
}

func Test_calculate_all_zero_values(t *testing.T) {
	checkSum := bank_ocr.CalculateCheckSum(bank_ocr.MustAccount(0, 0, 0, 0, 0, 0, 0, 0, 0))

	assert.Equal(t, 0, checkSum)
}

func Test_calculate_d1_coefficient_value_2(t *testing.T) {
	checkSum := bank_ocr.CalculateCheckSum(bank_ocr.MustAccount(0, 0, 0, 0, 0, 0, 0, 0, 2))

	assert.Equal(t, 2, checkSum)
}

func Test_calculate_d1_d9_coefficient_value_1(t *testing.T) {
	checkSum := bank_ocr.CalculateCheckSum(bank_ocr.MustAccount(1, 0, 0, 0, 0, 0, 0, 0, 1))

	assert.Equal(t, 10, checkSum)
}

func Test_calculate_d1_d9_coefficient_value_2(t *testing.T) {
	checkSum := bank_ocr.CalculateCheckSum(bank_ocr.MustAccount(2, 0, 0, 0, 0, 0, 0, 0, 2))

	assert.Equal(t, 20, checkSum)
}

func Test_calculate_d1_d9_coefficient_different_values(t *testing.T) {
	checkSum := bank_ocr.CalculateCheckSum(bank_ocr.MustAccount(1, 0, 0, 0, 0, 0, 0, 0, 2))

	assert.Equal(t, 11, checkSum)
}

func Test_calculate_d1_d9_coefficient_different_values_v2(t *testing.T) {
	checkSum := bank_ocr.CalculateCheckSum(bank_ocr.MustAccount(9, 0, 0, 0, 0, 0, 0, 0, 9))

	assert.Equal(t, 90, checkSum)
}

func Test_calculate_d1_d8_coefficient_different_values_v2(t *testing.T) {
	checkSum := bank_ocr.CalculateCheckSum(bank_ocr.MustAccount(0, 1, 0, 0, 0, 0, 0, 0, 1))

	assert.Equal(t, 9, checkSum)
}

func Test_calculate_d1_d7_coefficient_different_values_v2(t *testing.T) {
	checkSum := bank_ocr.CalculateCheckSum(bank_ocr.MustAccount(0, 0, 1, 0, 0, 0, 0, 0, 1))

	assert.Equal(t, 8, checkSum)
}

func Test_calculate_all_coefficient_has_values(t *testing.T) {
	checkSum := bank_ocr.CalculateCheckSum(bank_ocr.MustAccount(1, 1, 1, 1, 1, 1, 1, 1, 1))

	assert.Equal(t, 45, checkSum)
}

func Test_digits_from_zero_to_eight(t *testing.T) {
	assert.NotPanics(t, func() {
		bank_ocr.MustAccount(0, 1, 2, 3, 4, 5, 6, 7, 8)
	})
}

func Test_digits_from_one_to_nine(t *testing.T) {
	assert.NotPanics(t, func() {
		bank_ocr.MustAccount(1, 2, 3, 4, 5, 6, 7, 8, 9)
	})
}

func Test_digits_not_allowed_account_digit_10(t *testing.T) {
	assert.Panics(t, func() {
		bank_ocr.MustAccount(1, 2, 3, 4, 5, 6, 7, 8, 10)
	})
}

func Test_not_allowed_account_with_more_than_9_digits(t *testing.T) {
	assert.Panics(t, func() {
		bank_ocr.MustAccount(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	})
}

func Test_check_sum_is_valid(t *testing.T) {
	acc := bank_ocr.MustAccount(3, 4, 5, 8, 8, 2, 8, 6, 5)
	assert.True(t, bank_ocr.CheckSumIsValid(acc))
}

func Test_check_sum_is_invalid(t *testing.T) {
	acc := bank_ocr.MustAccount(3, 4, 5, 8, 8, 2, 8, 6, 3)
	assert.False(t, bank_ocr.CheckSumIsValid(acc))
}
