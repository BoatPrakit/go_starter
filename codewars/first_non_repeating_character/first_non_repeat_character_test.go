package codewars

import "testing"

func TestFirstNonRepeating(t *testing.T) {
	t.Run("" , func(t *testing.T) {
		want := ""
		result := FirstNonRepeating("")

		if result != want {
			t.Errorf("want %v but got %v", want, result)
		}
	})
}
