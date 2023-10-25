package fizzbuzz

import (
	"strconv"
	"strings"
)

type FizzBuzzLines []string

func (l FizzBuzzLines) Line(num int64) string {
	return l[num-1]
}

func FizzBuzz() FizzBuzzLines {
	lines := make(FizzBuzzLines, 100)

	for i := range lines {
		lines[i] = lineToStr(i)
	}

	return lines
}

func lineToStr(i int) string {
	strValue := strconv.Itoa(i + 1)
	acc := ""
	if (i+1)%3 == 0 {
		acc += "Fizz"
	}

	if strings.Contains(strValue, "3") {
		acc += "Fizz"
	}

	if (i+1)%5 == 0 {
		acc += "Buzz"
	}

	if strings.Contains(strValue, "5") {
		acc += "Buzz"
	}

	if acc == "" {
		return strValue
	}

	return acc
}
