package main

type User struct {
	Id   int64 `xorm:"pk autoincr"`
	Name string
}

type Account struct {
	Id     int64
	UserId int64 `xorm: "index"`
	Amount int64
}

type Car struct {
	Id     int64
	UserId int64 `xorm: "index"`
	Type   int
}
