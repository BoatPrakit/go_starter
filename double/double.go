package double

import "errors"

var (
	ErrMissingArgs   = errors.New("missing Arguments")
	ErrNoPersonFound = errors.New("no Person found")
)

type Person struct {
	FirstName   string
	LastName    string
	PhoneNumber string
}

type Queryer interface {
	Search(p []*Person, fn string, ln string) *Person
}

type PhoneBook struct {
	People []*Person
}

func (p *PhoneBook) Find(q Queryer, firstName string, lastName string) (string, error) {
	if firstName == "" || lastName == "" {
		return "", ErrMissingArgs
	}

	person := q.Search(p.People, firstName, lastName)

	if person == nil {
		return "", ErrNoPersonFound
	}

	return person.PhoneNumber, nil
}
