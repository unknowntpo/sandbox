package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

const (
	wantString = "Hello world from a temp file."
)

func TestCat(t *testing.T) {
	b := &bytes.Buffer{}
	b.Write([]byte(wantString))

	// create tmpfile using ioutil.TempFile
	tmpFile, cleanUpFunc := createTempFile(t)
	defer cleanUpFunc()

	// call Cat() read the content to buffer
	cat(tmpFile.Name(), b)
	// assert content
	got := b.String()
	want := wantString
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}

func createTempFile(t *testing.T) (*os.File, func()) {
	t.Helper()
	f, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatalf("Fail to create temp file: %v", err)
	}
	cleanUpFunc := func() {
		f.Close()
		os.Remove(f.Name())
	}
	return f, cleanUpFunc

}
