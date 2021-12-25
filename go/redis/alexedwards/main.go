package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type album struct {
	id            int
	title, artist string
	price         float64
	likes         int
}

func main() {
	// (2, "Electric Ladyland", "Jimi Hendrix", 4.95, 8
	err := addAlbum(&album{
		id:     2,
		title:  "Electric Ladyland",
		artist: "Jimi Hendrix",
		price:  4.95,
		likes:  8,
	})
	if err != nil {
		log.Fatal(err)
	}

	a, err := getAlbum(2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s by %s: £%.2f [%d likes]\n", a.title, a.artist, a.price, a.likes)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func addAlbum(a *album) error {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return fmt.Errorf("addAlbum: %v", err)
	}
	defer conn.Close()

	_, err = conn.Do(
		"HMSET",
		fmt.Sprintf("album:%d", a.id),
		"title", a.title,
		"artist", a.artist,
		"price", a.price,
		"likes", a.likes,
	)
	if err != nil {
		return fmt.Errorf("addAlbum: %v", err)
	}

	log.Println(a.title, " added!")
	return nil
}

func getAlbum(id int) (*album, error) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return nil, fmt.Errorf("failed to dial up redis server: %v", err)
	}
	defer conn.Close()

	// Issue a HGET command to retrieve the title for a specific album,
	// and use the Str() helper method to convert the reply to a string.
	title, err := redis.String(conn.Do("HGET", fmt.Sprintf("album:%d", id), "title"))
	if err != nil {
		return nil, fmt.Errorf("failed to get title of album %d: %v", id, err)
	}

	// Similarly, get the artist and convert it to a string.
	artist, err := redis.String(conn.Do("HGET", fmt.Sprintf("album:%d", id), "artist"))
	if err != nil {
		return nil, fmt.Errorf("failed to get artist of album %d: %v", id, err)
	}
	// And the price as a float64...
	price, err := redis.Float64(conn.Do("HGET", fmt.Sprintf("album:%d", id), "price"))
	if err != nil {
		return nil, fmt.Errorf("failed to get price of album %d: %v", id, err)
	}

	// And the number of likes as an integer.
	likes, err := redis.Int(conn.Do("HGET", fmt.Sprintf("album:%d", id), "likes"))
	if err != nil {
		return nil, fmt.Errorf("failed to get likes of album %d: %v", id, err)
	}

	return &album{
		id:     id,
		title:  title,
		artist: artist,
		price:  price,
		likes:  likes,
	}, nil
}
