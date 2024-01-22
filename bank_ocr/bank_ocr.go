package bank_ocr

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	numberOne = "" +
		"   " +
		"  |" +
		"  |"
	numberTwo = "" +
		" _ " +
		" _|" +
		"|_ "
	numberThree = "" +
		" _ " +
		" _|" +
		" _|"
	numberFour = "" +
		"   " +
		"|_|" +
		"  |"
	numberFive = "" +
		" _ " +
		"|_ " +
		" _|"
	numberSix = "" +
		" _ " +
		"|_ " +
		"|_|"
	numberSeven = "" +
		" _ " +
		"  |" +
		"  |"
	numberEight = "" +
		" _ " +
		"|_|" +
		"|_|"
	numberNine = "" +
		" _ " +
		"|_|" +
		" _|"
)

var numbers = map[string]int{
	numberOne:   1,
	numberTwo:   2,
	numberThree: 3,
	numberFour:  4,
	numberFive:  5,
	numberSix:   6,
	numberSeven: 7,
	numberEight: 8,
	numberNine:  9,
}

type digit string
type Account struct {
	Value [9]digit
}

func (a Account) Validate() string {
	if !a.CheckSumIsValid() {
		return "ERR"
	}
	return ""
}

func MustAccount(numbers ...int) Account {
	if len(numbers) > 9 {
		panic("numbers must be less than 9")
	}

	var account Account
	for i, n := range numbers {
		if n < 0 || n > 9 {
			panic(fmt.Sprintf("given digit %v Value out of range [0,9] in position %v", n, i))
		}

		digitStr := strconv.FormatInt(int64(n), 10)
		account.Value[i] = digit(digitStr)
	}

	return account
}

func ParseNumbers(s string) []Account {
	entries := ParseLines(s)
	entriesCount := len(entries)
	result := make([]Account, 0, entriesCount)

	for _, entry := range entries {
		result = append(result, parseLine(entry))
	}

	return result
}

func ParseLines(s string) []string {
	entries := strings.Split(s, "\n\n")
	return entries[:len(entries)-1]
}

func parseLine(entry string) Account {
	result := make([]int, 9)

	digits := parseDigit(entry)
	for i, digit := range digits {
		result[i] = numbers[digit]
	}

	return MustAccount(result...)
}

func parseDigit(entry string) []string {
	result := make([]string, 9)
	entryLines := strings.Split(entry, "\n")
	index := 0
	for i := 0; i < 9; i++ {
		numberStr := entryLines[0][index:index+3] + entryLines[1][index:index+3] + entryLines[2][index:index+3]
		result[i] = numberStr
		index += 3
	}
	return result
}

func CalculateCheckSum(account Account) int {
	acc := 0

	for i := 0; i < 9; i++ {
		digit, _ := strconv.ParseInt(string(account.Value[i]), 10, 32)

		acc += (9 - i) * int(digit)
	}

	return acc
}

func (a Account) CheckSumIsValid() bool {
	return CalculateCheckSum(a)%11 == 0
}
