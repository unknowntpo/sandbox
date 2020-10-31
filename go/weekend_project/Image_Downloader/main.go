package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	// Specify the folder to save the downloaded file
	download = "download"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func downloadFile(URL, fileName string) error {
	// Get the response bytes from URL
	// TODO: type of res? // a struct
	res, err := http.Get(URL)
	check(err)
	defer res.Body.Close()

	// Check the status code
	if res.StatusCode != 200 {
		// Diff between errors package and log ?
		return errors.New("Receive non 200 response code")
	}

	// Create folder download if it's not exist
	// Ref: stackoverflow:
	// 	- [mkdir if not exists using golang](https://stackoverflow.com/a/37932674)
	if _, err := os.Stat(download); os.IsNotExist(err) {
		os.Mkdir(download, 0755)
	}
	// Create an empty file at directory './download'
	path := "./" + download + "/" + fileName
	file, err := os.Create(path)
	check(err)
	defer file.Close()

	// Writes the bytes to the file
	// TODO: Implement image filter by bitwise operation
	_, err = io.Copy(file, res.Body)
	check(err)

	return nil
}
func main() {
	// TODO: Use flags package to handle command line argument
	if len(os.Args) != 4 {
		fmt.Println("Usage: ./Image_Downloader <URL> -o <output file name>")
		os.Exit(1)
	}
	// ./ImageDownloaer <URL> -o <output>
	fileName := os.Args[3]
	URL := os.Args[1]
	err := downloadFile(URL, fileName)
	check(err)
	fmt.Printf("File %s is downloaded \n", fileName)

}
