package main

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type database struct {
	users map[string]*user
}

type user struct {
	name         string
	passwordHash []byte
}

func main() {
	// signup
	db := &database{
		users: make(map[string]*user),
	}
	err := db.signUp("eric", "pa55word")
	if err != nil {
		fmt.Println(err)
	}

	// signin
	err = db.signIn("eric", "pa55word")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(db)
}

func (d *database) signUp(name, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return fmt.Errorf("fail to generate hash: %v", err)
	}

	u := &user{name: name, passwordHash: hash}
	// check if user exist or not
	if _, ok := d.users[name]; ok {
		return errors.New("user name exist")
	}

	d.users[name] = u

	fmt.Println("signup successfully!")

	return nil
}

func (d *database) signIn(name, password string) error {
	u := &user{}

	// Get user from database.
	if _, ok := d.users[name]; !ok {
		return errors.New("user does not exist")
	} else {
		u = d.users[name]
	}

	err := bcrypt.CompareHashAndPassword(u.passwordHash, []byte(password))
	if err != nil {
		return fmt.Errorf("fail to login: %v", err)
	}

	fmt.Println("signin successfully!")
	return nil
}
