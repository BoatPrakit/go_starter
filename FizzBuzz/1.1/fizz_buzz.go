package fizzbuzz

import "strconv"

func FizzBuzz(input int) string {
	if isFizz(input) && input%5 == 0 {
		return "FizzBuzz"
	}

	if isFizz(input) {
		return "Fizz"
	}

	if input%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(input)
}

func isFizz(input int) bool {
	return input%3 == 0
}
