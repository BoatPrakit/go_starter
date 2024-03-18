package book_seller

import "testing"

func TestBookSeller(t *testing.T) {
	t.Run("Should return empty string when books and categories are empty", func(t *testing.T) {
		books := []string{}
		categories := []string{}
		want := ""

		got := StockList(books, categories)

		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})

	t.Run("Should return A: 500 when books has A", func(t *testing.T) {
		books := []string{"A 100", "A 200", "A 200"}
		categories := []string{"A"}
		want := "(A : 500)"

		got := StockList(books, categories)

		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})

	t.Run("Should ", func(t *testing.T) {
		books := []string{"A 100", "B 200", "A 200"}
		categories := []string{"A"}
		want := "(A : 300)"

		got := StockList(books, categories)

		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})

	t.Run("Should ", func(t *testing.T) {
		books := []string{"ABB 100", "B 200", "ACC 200"}
		categories := []string{"A"}
		want := "(A : 300)"

		got := StockList(books, categories)

		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})

	t.Run("Should ", func(t *testing.T) {
		books := []string{"ABB 100", "B 200", "ACC 200"}
		categories := []string{"B"}
		want := "(B : 200)"

		got := StockList(books, categories)

		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})

	t.Run("Should ", func(t *testing.T) {
		books := []string{"ABB 100", "B 200", "ACC 200"}
		categories := []string{"A", "B"}
		want := "(A : 300) - (B : 200)"

		got := StockList(books, categories)

		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})
	t.Run("Should ", func(t *testing.T) {
		books := []string{"ABB 100", "B 200", "ACC 200"}
		categories := []string{"A", "B"}
		want := "(A : 300) - (B : 200)"

		got := StockList(books, categories)

		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})

	t.Run("Should ", func(t *testing.T) {
		books := []string{"ABB 100", "B 200", "ACC 200", "DDD 100"}
		categories := []string{"A", "B", "D"}
		want := "(A : 300) - (B : 200) - (D : 100)"

		got := StockList(books, categories)

		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})
}
