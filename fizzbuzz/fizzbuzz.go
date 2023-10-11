package fizzbuzz

import "strconv"

func FizzBuzz() []string {
	lines := make([]string, 100)
	for i, _ := range lines {
		lines[i] = strconv.Itoa(i + 1)
		if (i+1)%3 == 0 {
			lines[i] = "Fizz"
		}
	}

	return lines
}
