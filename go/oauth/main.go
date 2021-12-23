package main

import (
	"fmt"
	"io"
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
	q.Add("redirect_uri", fmt.Sprintf("http://localhost:%d/callback", port))
	q.Add("scope", "user, repo")
	q.Add("allow_signup", "true")
	rawQuery := q.Encode()
	//fmt.Println(rawQuery)
	// TODO: redirect user to that url
	redirURL += "?" + rawQuery
	fmt.Println(redirURL)
	http.Redirect(w, r, redirURL, http.StatusSeeOther)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("hello"))

	// extract code from r params
	// send request to get access token
	// POST https://github.com/login/oauth/access_token
	// params:
	// client_id
	// client_secret
	// code
	// redirect
	code := r.URL.Query().Get("code")
	if code == "" {
		fmt.Fprintf(w, "failed to get code from url parameter")
		return
	}
	fmt.Fprintf(w, "code: %s", code)

	client := http.Client{}

	r, err := http.NewRequest(http.MethodPost, "https://github.com/login/oauth/access_token", nil)
	if err != nil {
		fmt.Fprintf(w, "failed to send request: %v", err)
		return
	}
	q := url.Values{}
	q.Add("code", code)
	q.Add("client_id", CLIENT_ID)
	q.Add("client_secret", CLIENT_SECRET)
	q.Add("redirect_uri", "http://localhost:4444/callback")

	r.URL.RawQuery = q.Encode()

	r.Header.Set("Accept", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		fmt.Fprintf(w, "failed to send request: %v", err)
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "failed to read resp body: %v", err)
		return
	}
	fmt.Println("response: \n", string(b))
	// TODO: parse response body and access github user repo
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
