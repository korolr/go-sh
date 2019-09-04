package main

import (
	"github.com/korolr/go-sh/db"
	"github.com/korolr/go-sh/model"
)

func main() {
	d, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer d.Close()

	d.DropTableIfExists(&model.User{})
	d.CreateTable(&model.User{})
}
