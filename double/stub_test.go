package double

import "testing"

type StubSearcher struct {
	phone string
}

func (s *StubSearcher) Search(p []*Person, firstName string, lastName string) *Person {
	return &Person{
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "+66 123456789",
	}
}

func TestFoundPhoneNumber(t *testing.T) {
	phoneBook := &PhoneBook{}

	want := "+66 123456789"

	phoneNumber, err := phoneBook.Find(&StubSearcher{want}, "John", "Doe")

	if err != nil {
		t.Errorf("should not be error %v", err)
	}

	if phoneNumber != want {
		t.Errorf("want %v but got %v", want, phoneNumber)
	}
}
