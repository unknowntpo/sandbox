package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	must(addAlbum(2, "Electric Ladyland", "Jimi Hendrix", 4.95, 8))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func addAlbum(id int, title, artist string, price float64, likes int) error {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return fmt.Errorf("addAlbum: %v", err)
	}
	defer conn.Close()

	_, err = conn.Do(
		"HMSET",
		fmt.Sprintf("album:%d", id),
		"title", title,
		"artist", artist,
		"price", price,
		"likes", likes,
	)
	if err != nil {
		return fmt.Errorf("addAlbum: %v", err)
	}

	log.Println(title, " added!")
	return nil
}
