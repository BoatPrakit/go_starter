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
}
