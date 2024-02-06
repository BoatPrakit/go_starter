package fizzbuzz

import "strconv"

func FizzBuzz(input int) string {
	if isFizz(input) && isBuzz(input) {
		return "FizzBuzz"
	}

	if isFizz(input) {
		return "Fizz"
	}

	if isBuzz(input) {
		return "Buzz"
	}

	return strconv.Itoa(input)
}

func isFizz(input int) bool {
	return input%3 == 0
}

func isBuzz(input int) bool {
	return input%5 == 0
}
