package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("should be return 1 when input 1", func(t *testing.T) {
		input := 1
		want := "1"

		result := FizzBuzz(1)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})

	t.Run("should be return 2 when input 2", func(t *testing.T) {
		input := 2
		want := "2"

		result := FizzBuzz(2)

		if result != want {
			t.Errorf("FizzBuzz should be return %v when input %v", want, input)
		}
	})
}
