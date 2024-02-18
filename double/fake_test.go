package double

import "testing"

type FakeSearcher struct{}

func (f *FakeSearcher) Search(p []*Person, fn string, ln string) *Person {
	if len(p) == 0 {
		return nil
	}

	return p[0]
}

func TestPhoneBookReturnEmptyString(t *testing.T) {

	fakeSearch := FakeSearcher{}
	phoneBook := PhoneBook{}

	phoneNumber, _ := phoneBook.Find(&fakeSearch, "aaa", "bbb")

	if phoneNumber != "" {
		t.Errorf("Should return empty string but got %v", phoneNumber)
	}
}

func TestPhoneBookReturnNumber(t *testing.T) {

	want := "+66 12345678"
	fakeSearch := FakeSearcher{}
	phoneBook := PhoneBook{
		People: []*Person{&Person{PhoneNumber: "+66 12345678"}},
	}

	phoneNumber, _ := phoneBook.Find(&fakeSearch, "aaa", "bbb")

	if phoneNumber == "" {
		t.Errorf("want %v but got %v", want, phoneNumber)
	}
}
