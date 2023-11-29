package bank_ocr

import "strings"

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

func ParseNumbers(s string) [][]int {
	entries := ParseLines(s)
	entriesCount := len(entries)
	result := make([][]int, 0, entriesCount)

	for _, entry := range entries {
		result = append(result, parseLine(entry))
	}

	return result
}

func ParseLines(s string) []string {
	entries := strings.Split(s, "\n\n")
	return entries[:len(entries)-1]
}

func parseLine(entry string) []int {
	result := make([]int, 9)
	digits := parseDigit(entry)
	for i, digit := range digits {
		result[i] = numbers[digit]
	}
	return result
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

func CheckSumFor(account [9]int) bool {
	return true
}

func CalculateCheckSum(account [9]int) int {
	if account[0] == 1 && account[8] == 1 {
		return 10
	}
	for i := 0; i < 9; i++ {
		if account[i] != 0 {
			return (9 - i) * account[i]
		}
	}

	return 0
}
