package main

import "testing"

// Fake: implementations that look like the real thing. But they only do the trick in the scope of the test.
// e.g. in-memory db.
type FakeSearcher struct{}

func (fs FakeSearcher) Search(people []*Person, firstName string, lastName string) *Person {
	if len(people) == 0 {
		return nil
	}

	return people[0]
}

func TestFindCallsSearchAndReturnsEmptyStringForNoPerson(t *testing.T) {
	phonebook := &Phonebook{}
	fake := &FakeSearcher{}

	phone, _ := phonebook.Find(fake, "Jane", "Doe")

	if phone != "" {
		t.Errorf("Wanted '', got '%s'", phone)
	}
}
