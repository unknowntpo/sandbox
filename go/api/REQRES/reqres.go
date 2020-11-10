package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

// Get single user from reqres.in
const UserURL = "https://reqres.in/api/users/2"
const tmpl = `---------------------------------
ID: {{.Data.Id}}
Email: {{.Data.Email}}
Name: {{.Data.FirstName}} {{.Data.LastName}}
Avatar: {{.Data.Avatar}}
Company: {{.Ad.Company}}
Url: {{.Ad.Url}}
Text: {{.Ad.Text}}
`

type UserResult struct {
	Data *Data `json:"data"`
	Ad   *Ad   `json:"ad"`
}

type Data struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type Ad struct {
	Company string `json:"company"`
	Url     string `json:"url"`
	Text    string `json:"text"`
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

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func main() {
	result, err := SearchUser()
	if err != nil {
		log.Fatal(err)
	}
	//Debug output
	//fmt.Printf("%+v\n", result.Data)
	//fmt.Printf("%+v\n", result.Ad)
	t := template.Must(template.New("t1").Parse(tmpl))
	check(t.Execute(os.Stdout, result))
}
