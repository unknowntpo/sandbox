// Ref: https://gitea.com/xorm/xorm
package main

import (
	"log"
	"time"

	"xorm.io/xorm"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	log.Println("Engine started")

	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		log.Printf("failed to create new engine: %v", err)

	}

	log.Println("Engine started")

	// config the logs
	engine.ShowSQL(true)

	// Sync the table to database
	if err := engine.Sync(new(User)); err != nil {
		log.Printf("failed to sync table user: %v", err)
	}

	results, err := engine.Query("select * from user")
	if err != nil {
		log.Printf("failed to select * from user: %v", err)
	}

	log.Println(results)

	/*
		results, err = engine.QueryString("select * from user")
		must(err)

		results, err = engine.Where("a = 1").QueryString()
		must(err)

		results, err = engine.QueryInterface("select * from user")
		must(err)

		results, err = engine.Where("a = 1").QueryInterface()
		must(err)
	*/

	defer engine.Close()
}
