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
}
