package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("should be return 1 when input 1", func(t *testing.T) {
		input := 1
		want := "1"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return 2 when input 2", func(t *testing.T) {
		input := 2
		want := "2"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return Fizz when input 3", func(t *testing.T) {
		input := 3
		want := "Fizz"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return 4 when input 4", func(t *testing.T) {
		input := 4
		want := "4"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return Buzz when input 5", func(t *testing.T) {
		input := 5
		want := "Buzz"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return Fizz when input 6", func(t *testing.T) {
		input := 6
		want := "Fizz"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return 7 when input 7", func(t *testing.T) {
		input := 7
		want := "7"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return Fizz when input 9", func(t *testing.T) {
		input := 9
		want := "Fizz"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return Buzz when input 10", func(t *testing.T) {
		input := 10
		want := "Buzz"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return Fizz when input 12", func(t *testing.T) {
		input := 12
		want := "Fizz"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return FizzBuzz when input 15", func(t *testing.T) {
		input := 15
		want := "FizzBuzz"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return Fizz when input 18", func(t *testing.T) {
		input := 18
		want := "Fizz"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return Buzz when input 20", func(t *testing.T) {
		input := 20
		want := "Buzz"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return FizzBuzz when input 30", func(t *testing.T) {
		input := 30
		want := "FizzBuzz"

		result := FizzBuzz(input)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})
}

func TestIsFizz(t *testing.T) {
	t.Run("should be return true when input 3", func(t *testing.T) {
		input := 3
		want := true

		result := isFizz(input)

		if result != want {
			t.Errorf("isFizz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return true when input 6", func(t *testing.T) {
		input := 6
		want := true

		result := isFizz(input)

		if result != want {
			t.Errorf("isFizz should be return %v when input %v", want, input)
		}
	})
}
