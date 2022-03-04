package main

import (
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"xorm.io/xorm"
)

func TestGetUser(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	assert.NoError(t, err)
	engine.Sync2(new(User), new(Account), new(Car))

	// Insert users
	users := []User{
		{Name: "eric"},
		{Name: "frank"},
		{Name: "angus"},
	}
	// insert 1 user
	user1 := User{Name: "happy"}
	_, err = engine.Insert(&user1)
	assert.NoError(t, err)
	t.Log(&user1) // got {27 happy}

	// insert users
	err = insertUsers(t, engine, &users)
	assert.NoError(t, err)
	t.Log(users)

	gotUsers := []User{}
	engine.Table("User").Find(&gotUsers)
	t.Log(gotUsers)
}

func insertUsers(t *testing.T, engine *xorm.Engine, users *[]User) error {
	_, err := engine.Insert(users)
	if err != nil {
		return fmt.Errorf("failed on insertUsers: %v", err)
	}
	return nil
}
