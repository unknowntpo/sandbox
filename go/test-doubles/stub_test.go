package main

import (
	"testing"
)

// When we want to make the Searcher implementation return an actual value that we can assert against,
// we need a stub. Instead of understanding what it would take to set up and/or implement a proper Searcher,
// you just create a stub implementation that returns only one value.
// This is the idea behind stubs.
type StubSearcher struct {
	phone string
}

func (ss StubSearcher) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

func TestFindReturnsPerson(t *testing.T) {
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}

	phone, _ := phonebook.Find(StubSearcher{fakePhone}, "Jane", "Doe")

	if phone != fakePhone {
		t.Errorf("Want '%s', got '%s'", fakePhone, phone)
	}
}