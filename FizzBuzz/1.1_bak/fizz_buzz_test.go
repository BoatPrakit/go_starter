package fizz_buzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("should be 1 when input 1", func(t *testing.T) {
		// Arrange
		input := 1
		want := "1"

		// Act
		got := FizzBuzz(input)

		// Assert
		if got != want {
			t.Errorf("FizzBuzz  should be return %v when input %v", want, input)
		}
	})

	t.Run("should be 2 when input 2", func(t *testing.T) {
		input := 2
		want := "2"

		got := FizzBuzz(input)

		if got != want {
			t.Errorf("FizzBuzz  should be return %v when input %v", want, input)
		}
	})

	t.Run("should be 3 when input 3", func(t *testing.T) {
		input := 3
		want := "Fizz"

		got := FizzBuzz(input)

		if got != want {
			t.Errorf("FizzBuzz  should be return %v when input %v", want, input)
		}
	})

	t.Run("should be 4 when input 4", func(t *testing.T) {
		input := 4
		want := "4"

		got := FizzBuzz(input)

		if got != want {
			t.Errorf("FizzBuzz  should be return %v when input %v", want, input)
		}
	})

	t.Run("should be 5 when input 5", func(t *testing.T) {
		input := 5
		want := "Buzz"

		got := FizzBuzz(input)

		if got != want {
			t.Errorf("FizzBuzz  should be return %v when input %v", want, input)
		}
	})

	t.Run("should be 6 when input 6", func(t *testing.T) {
		input := 6
		want := "Fizz"

		got := FizzBuzz(input)

		if got != want {
			t.Errorf("FizzBuzz  should be return %v when input %v", want, input)
		}
	})

	t.Run("should be 7 when input 7", func(t *testing.T) {
		input := 7
		want := "7"

		got := FizzBuzz(input)

		if got != want {
			t.Errorf("FizzBuzz  should be return %v when input %v", want, input)
		}
	})

	t.Run("should be 8 when input 8", func(t *testing.T) {
		input := 8
		want := "8"

		got := FizzBuzz(input)

		if got != want {
			t.Errorf("FizzBuzz  should be return %v when input %v", want, input)
		}
	})

	t.Run("should be 9 when input 9", func(t *testing.T) {
		input := 9
		want := "Fizz"

		got := FizzBuzz(input)

		if got != want {
			t.Errorf("FizzBuzz  should be return %v when input %v", want, input)
		}
	})

	t.Run("should be 10 when input 10", func(t *testing.T) {
		input := 10
		want := "Buzz"

		got := FizzBuzz(input)

		if got != want {
			t.Errorf("FizzBuzz  should be return %v when input %v", want, input)
		}
	})

	t.Run("should be 15 when input 15", func(t *testing.T) {
		input := 15
		want := "FizzBuzz"

		got := FizzBuzz(input)

		if got != want {
			t.Errorf("FizzBuzz  should be return %v when input %v", want, input)
		}
	})
}
