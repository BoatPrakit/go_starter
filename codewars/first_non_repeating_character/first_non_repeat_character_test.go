package codewars

import "testing"

func TestFirstNonRepeating(t *testing.T) {
	t.Run("should return empty string when input empty string", func(t *testing.T) {
		want := ""
		result := FirstNonRepeating("")

		if result != want {
			t.Errorf("want %v but got %v", want, result)
		}
	})

	t.Run("should return empty string when input aaaaaa", func(t *testing.T) {
		want := ""
		input := "aaaaaa"
		result := FirstNonRepeating(input)

		if result != want {
			t.Errorf("want %v but got %v", want, result)
		}
	})

	t.Run("should return a when input abcdefg", func(t *testing.T) {
		want := "a"
		input := "abcdefg"
		result := FirstNonRepeating(input)

		if result != want {
			t.Errorf("want %v but got %v", want, result)
		}
	})

	t.Run("should return c when input abacddbe", func(t *testing.T) {
		want := "c"
		input := "abacddbe"
		result := FirstNonRepeating(input)

		if result != want {
			t.Errorf("want %v but got %v", want, result)
		}
	})
}
