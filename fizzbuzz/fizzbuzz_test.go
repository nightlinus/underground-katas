package fizzbuzz_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"underground-katas/fizzbuzz"
)

func Test_FizzBuzz_print_100_lines(t *testing.T) {
	assert.Equal(t, 100, len(fizzbuzz.FizzBuzz()))
}

func Test_FizzBuzz_each_line_prints_its_number(t *testing.T) {
	res := fizzbuzz.FizzBuzz()

	assert.Equal(t, res[0], "1")
	assert.Equal(t, res[1], "2")
	assert.Equal(t, res[99], "100")
}
