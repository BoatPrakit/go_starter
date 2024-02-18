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
	t.Run("should return T when input sTreSS", func(t *testing.T) {
		want := "T"
		input := "sTreSS"
		result := FirstNonRepeating(input)

		if result != want {
			t.Errorf("want %v but got %v", want, result)
		}
	})
	t.Run("should return H when input Hello! World!", func(t *testing.T) {
		want := "H"
		input := "Hello! World!"
		result := FirstNonRepeating(input)

		if result != want {
			t.Errorf("want %v but got %v", want, result)
		}
	})
	t.Run("should return a when input a", func(t *testing.T) {
		want := "a"
		input := "a"
		result := FirstNonRepeating(input)

		if result != want {
			t.Errorf("want %v but got %v", want, result)
		}
	})
}
