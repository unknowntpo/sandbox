package main

import (
	"testing"
)

// DummySearcher is a dummy instance for searcher interface,
// all we need is an instance that we can just pass in, and the compiler won’t complain.
type DummySearcher struct{}

func (ds DummySearcher) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{}
}

func TestFindReturnsError(t *testing.T) {
	phonebook := &Phonebook{}

	want := ErrMissingArgs
	_, got := phonebook.Find(DummySearcher{}, "", "")

	if got != want {
		t.Errorf("Want '%s', got '%s'", want, got)
	}
}

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

// Spy: funtionality of stub + keep track of method invocation.
// That’s what spy is - a stub that keeps track of invocations of its methods.
type SpySearcher struct {
	phone           string
	searchWasCalled bool
}

func (ss *SpySearcher) Search(people []*Person, firstName, lastName string) *Person {
	ss.searchWasCalled = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

func TestFindCallsSearchAndReturnsPerson(t *testing.T) {
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}
	spy := &SpySearcher{phone: fakePhone}

	phone, _ := phonebook.Find(spy, "Jane", "Doe")

	if !spy.searchWasCalled {
		t.Errorf("Expected to call 'Search' in 'Find', but it wasn't.")
	}

	if phone != fakePhone {
		t.Errorf("Want '%s', got '%s'", fakePhone, phone)
	}
}
