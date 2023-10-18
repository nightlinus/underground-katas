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
		strValue := strconv.Itoa(i + 1)
		if (i+1)%3 == 0 {
			lines[i] = "Fizz"
		}

		if (i+1)%5 == 0 {
			lines[i] = "Buzz"
		}

		if (i+1)%3 == 0 && (i+1)%5 == 0 {
			lines[i] = "FizzBuzz"
		}

		if strings.Contains(strValue, "3") {
			lines[i] = lines[i] + "Fizz"
		}

		if lines[i] == "" {
			lines[i] = strValue
		}
	}

	return lines
}
