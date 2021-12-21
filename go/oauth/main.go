package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	CLIENT_ID         = os.Getenv("GITHUB_OAUTH_CLIENT_ID")
	CLIENT_SECRET     = os.Getenv("GITHUB_OAUTH_CLIENT_SECRET")
	port          int = 4444
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
	/*
		GET https://github.com/login/oauth/authorize
		client_id= ...
		redirect_uri=localhost:4444/callback
		scope=user repo
		allow_signup=true
	*/
	redirURL := "https://github.com/login/oauth/authorize"
	_ = redirURL
	q := url.Values{}
	q.Add("client_id", fmt.Sprint(CLIENT_ID))
	q.Add("redirect_uri", fmt.Sprintf("http://localhost:%d", port))
	q.Add("scope", "user, repo")
	q.Add("allow_signup", "true")
	rawQuery := q.Encode()
	fmt.Println(rawQuery)
	// TODO: redirect user to that url
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
	// send request to get access token
	// POST https://github.com/login/oauth/access_token
	// params:
	// client_id
	// client_secret
	// code
	// redirect

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
	// TOOD: Display user information get from github
	/*
		Authorization: token OAUTH-TOKEN
		GET https://api.github.com/user
	*/
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/callback", callbackHandler)
	mux.HandleFunc("/", homeHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("oauth example app listen on :%d", port)
	log.Fatal(srv.ListenAndServe())
}
