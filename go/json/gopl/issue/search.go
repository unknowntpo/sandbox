package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues queries the Github issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	// What does query do ?
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Println("URL with query: ", IssuesURL+"?q="+q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	//
	/*
		req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set(
			"Accept", "application/vnd.github.v3.text-match+json")
		resp, err := http.DefaultClient.Do(req)
	*/
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
