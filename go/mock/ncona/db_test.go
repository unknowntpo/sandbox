package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"errors"
	"testing"
)

type MockDatabase struct {
	mock.Mock
}

func (db *MockDatabase) connect() error {
	args := db.Called()
	return args.Error(0)
}

func (db *MockDatabase) sendMessage(message *string) error {
	args := db.Called(message)
	return args.Error(0)
}

func TestSuccess(t *testing.T) {
	db := new(MockDatabase)
	message := "Hello"

	// Set expectations
	db.On("connect").Return(nil)
	db.On("sendMessage", &message).Return(nil)

	err := Talk(db, &message)

	assert.Equal(t, nil, err, "No error")
	db.AssertExpectations(t)
}

func TestErrorOnMessage(t *testing.T) {
	db := new(MockDatabase)
	message := "Hello"

	// Set expetations
	db.On("connect").Return(nil)
	db.On("sendMessage", &message).Return(errors.New("Some error"))

	err := Talk(db, &message)

	assert.NotEqual(t, nil, err, "An error is thrown if sendMessage fails")
	db.AssertExpectations(t)
}
