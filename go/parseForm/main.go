package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func handleEmail(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	details := ContactDetails{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	// Do something with details
	// Display message received
	fmt.Println("Get message:")
	fmt.Println("Email:", details.Email)
	fmt.Println("Subject:", details.Subject)
	fmt.Println("Message:", details.Message)

	tmpl.Execute(w, struct{ Success bool }{true})

}
func main() {

	http.HandleFunc("/", handleEmail)
	http.ListenAndServe(":8080", nil)
}
