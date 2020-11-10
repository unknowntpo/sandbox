package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Get single user from reqres.in
const UserURL = "https://reqres.in/api/users/2"

type UserResult struct {
	Data *Data `json:"data"`
	Ad   *Ad   `json:"ad"`
}

type Data struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	FirstNname string `json:"first_name"`
	LastNname  string `json:"last_name"`
	Avatar     string `json:"avatar"`
}

type Ad struct {
	Company string
	Url     string
	Text    string
}

func SearchUser() (*UserResult, error) {
	fmt.Println("URL: ", UserURL)
	resp, err := http.Get(UserURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Fail to get resource from URL")
	}

	var result UserResult
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {
	result, err := SearchUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result.Data)
	fmt.Printf("%+v\n", result.Ad)
}
