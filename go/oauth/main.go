package main

import (
	"encoding/json"
	"errors"
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

type TokenInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}
type application struct {
	tokenInfo *TokenInfo
}

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
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

func (app *application) callbackHandler(w http.ResponseWriter, r *http.Request) {
	tInfo, err := getTokenInfo(r)
	if err != nil {
		log.Printf("failed to get token information from github callback redirection: %v", err)
		return
	}
	app.tokenInfo = tInfo
	redirURL := fmt.Sprintf("http://localhost:%d/", port)
	http.Redirect(w, r, redirURL, http.StatusSeeOther)
}

// getAccessToken returns the access token of specific user.
func getTokenInfo(r *http.Request) (tInfo *TokenInfo, err error) {
	code := r.URL.Query().Get("code")
	if code == "" {
		err = errors.New("failed to get code from url parameter")
	}

	client := http.Client{}

	r, err = http.NewRequest(http.MethodPost, "https://github.com/login/oauth/access_token", nil)
	if err != nil {
		err = fmt.Errorf("failed to send request: %v", err)
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
		err = fmt.Errorf("failed to send request: %v", err)
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed to read resp body: %v", err)
		return
	}
	fmt.Println("response: \n", string(b))
	// TODO: parse response body and access github user repo

	// {"access_token":"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","token_type":"bearer","scope":"repo,user"}
	tInfo = &TokenInfo{}
	err = json.Unmarshal(b, tInfo)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal response from github: %v", err)
		return
	}
	log.Printf("%+v", tInfo)

	return
}

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
	// TODO: Display user information get from github based on app.tokenInfo
	/*
		Authorization: token OAUTH-TOKEN
		GET https://api.github.com/user
	*/
}

func main() {
	app := &application{}
	mux := http.NewServeMux()
	mux.HandleFunc("/login", app.loginHandler)
	mux.HandleFunc("/callback", app.callbackHandler)
	mux.HandleFunc("/", app.homeHandler)

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
