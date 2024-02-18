package double

import "testing"

type DummySeacher struct{}

func (d *DummySeacher) Search(p []Person, firstName string, lastName string) *Person {
	return &Person{}
}

func TestNoArgFound(t *testing.T) {
	phoneBook := &PhoneBook{}

	want := ErrMissingArgs
	_, err := phoneBook.Find(&DummySeacher{}, "", "")

	if err != want {
		t.Errorf("want %v but got %v", want, err)
	}
}
