package main

import (
	"testing"
)

// Mocks work by setting certain expectations on them.
// We implement some stub-like functionality while keeping track of the methods that have been called.
// Finally, we ask the mock to verify if our code met all of its expectations.
type MockSearcher struct {
	phone         string
	methodsToCall map[string]bool
}

func (ms *MockSearcher) Search(people []*Person, firstName, lastName string) *Person {
	ms.methodsToCall["Search"] = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ms.phone,
	}
}

func (ms *MockSearcher) ExpectToCall(methodName string) {
	if ms.methodsToCall == nil {
		ms.methodsToCall = make(map[string]bool)
	}
	ms.methodsToCall[methodName] = false
}

func (ms *MockSearcher) Verify(t *testing.T) {
	for methodName, called := range ms.methodsToCall {
		if !called {
			t.Errorf("Expected to call '%s', but it wasn't.", methodName)
		}
	}
}

func TestFindCallsSearchAndReturnsPersonUsingMock(t *testing.T) {
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}
	mock := &MockSearcher{phone: fakePhone}
	mock.ExpectToCall("Search")

	phone, _ := phonebook.Find(mock, "Jane", "Doe")

	if phone != fakePhone {
		t.Errorf("Want '%s', got '%s'", fakePhone, phone)
	}

	mock.Verify(t)
}
