package bank_ocr

import "strings"

const (
	oneNumberEntry = `                           
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |`
	twoNumberEntry = `_  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
|_ |_ |_ |_ |_ |_ |_ |_ |_ `
	mixedNumbersEntry = `   _  _  _  _  _  _  _  _ 
 |  _| _| _| _| _| _| _| _|
 | |_ |_ |_ |_ |_ |_ |_ |_ `
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
)

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
	entryLines := strings.Split(entry, "\n")
	index := 0
	for i := 0; i < 9; i++ {
		number := entryLines[0][index:index+3] + entryLines[1][index:index+3] + entryLines[2][index:index+3]
		if number == numberOne {
			result[i] = 1
		}

		if number == numberTwo {
			result[i] = 2
		}
		if number == numberThree {
			result[i] = 3
		}
		index += 3
	}
	return result
}
