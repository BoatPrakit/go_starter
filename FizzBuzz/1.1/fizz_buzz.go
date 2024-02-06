package fizzbuzz

import "strconv"

func FizzBuzz(input int) string {
	if input == 3 {
		return "Fizz"
	}

	return strconv.Itoa(input)
}
