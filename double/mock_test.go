package double

import "testing"

type MockSearcher struct {
	methodCalled map[string]bool
	phoneNumber  string
}

func (m *MockSearcher) Search(p []*Person, fn string, ln string) *Person {
	m.methodCalled["Search"] = true

	return &Person{
		FirstName:   fn,
		LastName:    ln,
		PhoneNumber: m.phoneNumber,
	}
}

func (m *MockSearcher) ExpectToCalled(methodName string) {
	if m.methodCalled == nil {
		m.methodCalled = make(map[string]bool)
	}

	m.methodCalled[methodName] = false
}

func (m *MockSearcher) Verify(t *testing.T) {
	for method, called := range m.methodCalled {
		if !called {
			t.Errorf("Method %v should been called but it does not", method)
		}
	}
}

func TestMethodSearchShouldBeenCalledAndReturnPerson(t *testing.T) {
	mock := &MockSearcher{
		phoneNumber: "+66 12345678",
	}
	phoneBook := &PhoneBook{}
	mock.ExpectToCalled("Search")

	want := "+66 12345678"
	phoneNumber, err := phoneBook.Find(mock, "John", "Doe")

	if err != nil {
		t.Errorf("should not be error %v", err)
	}

	if phoneNumber != want {
		t.Errorf("want %v but got %v", want, phoneNumber)
	}

	mock.Verify(t)
}
