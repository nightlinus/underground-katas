package fizzbuzz

import "strconv"

type FizzBuzzLines []string

func (l FizzBuzzLines) Line(num int64) string {
	return l[num-1]
}

func FizzBuzz() FizzBuzzLines {
	lines := make(FizzBuzzLines, 100)
	for i, _ := range lines {
		lines[i] = strconv.Itoa(i + 1)
		if (i+1)%3 == 0 {
			lines[i] = "Fizz"
		}

		if (i+1)%5 == 0 {
			lines[i] = "Buzz"
		}
	}

	return lines
}
