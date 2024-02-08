package fizz_buzz

import "strconv"

func FizzBuzz(input int) string {
	if input%3 == 0 {
		return "Fizz"
	}

	if input == 5 {
		return "Buzz"
	}

	return strconv.Itoa(input)
}
