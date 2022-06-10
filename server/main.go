package main

import (
	"database/sql"
	"final-project/api"
	"final-project/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./tulisaja-app.db")
	if err != nil {
		panic(err)
	}

	usersRepo := repository.NewUserRepository(db)

	mainAPI := api.NewAPI(*usersRepo)
	mainAPI.Start()
}