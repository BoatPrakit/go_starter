package fizz_buzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("should be return 1 when input 1", func(t *testing.T) {
		input := 1
		want := "1"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should return %v when input %v", want, input)
		}
	})

	t.Run("should be return 2 when input 2", func(t *testing.T) {
		input := 2
		want := "2"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should return %v when input %v", want, input)
		}
	})

	t.Run("should be return Fizz when input 3", func(t *testing.T) {
		input := 3
		want := "Fizz"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should return %v when input %v", want, input)
		}
	})

	t.Run("should be return 4 when input 4", func(t *testing.T) {
		input := 4
		want := "4"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should return %v when input %v", want, input)
		}
	})
}
