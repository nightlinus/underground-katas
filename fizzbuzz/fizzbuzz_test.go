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

	assert.Equal(t, "1", res.Line(1))
	assert.Equal(t, "2", res.Line(2))
	assert.Equal(t, "98", res.Line(98))
}

func Test_line_number_devisable_by_3_print_Fizz(t *testing.T) {
	lines := fizzbuzz.FizzBuzz()

	assert.Equal(t, "Fizz", lines.Line(3))
	assert.Equal(t, "Fizz", lines.Line(6))
	assert.Equal(t, "Fizz", lines.Line(99))
}

func Test_line_number_devisable_by_5_print_Buzz(t *testing.T) {
	lines := fizzbuzz.FizzBuzz()

	assert.Equal(t, "Buzz", lines.Line(5))
	assert.Equal(t, "Buzz", lines.Line(10))
	assert.Equal(t, "Buzz", lines.Line(100))
}

func Test_line_number_devisable_by_5_and_3_print_FizzBuzz(t *testing.T) {
	lines := fizzbuzz.FizzBuzz()

	assert.Equal(t, "FizzBuzz", lines.Line(15))
	assert.Equal(t, "FizzBuzz", lines.Line(45))
	assert.Equal(t, "FizzBuzz", lines.Line(60))
	assert.Equal(t, "FizzBuzz", lines.Line(90))
}
