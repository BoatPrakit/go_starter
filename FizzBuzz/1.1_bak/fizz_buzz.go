package fizz_buzz

import "fmt"

func FizzBuzz(input int) string {
	fb := isFizzNa(input%3 == 0) + isBuzzNa(isBuzz(input)) + normal(!(input%3 == 0 || input%5 == 0), input)
	return fb
}

func normal(b bool, input int) string {
	m := map[bool]string{
		true:  fmt.Sprint(input),
		false: "",
	}
	return m[b]
}

// KISS => Keep It Simple and Stupid
// DRY => Don't Repeate Yourself
// YAGNI => You aren't Gona Need It

func isFizzNa(b bool) string {
	m := map[bool]string{
		true:  "Fizz",
		false: "",
	}
	return m[b]
}

func isBuzzNa(b bool) string {
	m := map[bool]string{
		true:  "Buzz",
		false: "",
	}
	return m[b]
}

func isFizz(input int) bool {
	return input%3 == 0
}

func isBuzz(input int) bool {
	return input%5 == 0
}
