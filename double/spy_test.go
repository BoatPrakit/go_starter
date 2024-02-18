package double

import "testing"

type SpySearcher struct {
	methodCalled bool
}

func (s *SpySearcher) Search(p []*Person, fn string, ln string) *Person {
	s.methodCalled = true

	return &Person{}
}

func TestMethodShouldBeCalled(t *testing.T) {
	want := true

	spy := &SpySearcher{}
	phoneBook := &PhoneBook{}

	_, err := phoneBook.Find(spy, "f", "s")

	if err != nil {
		t.Errorf("should not error %v", err)
	}

	if spy.methodCalled != want {
		t.Errorf("Expect to called Search in Find")
	}
}
