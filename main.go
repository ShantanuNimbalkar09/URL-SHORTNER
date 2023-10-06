package main

import (
	"app/database"
	"app/router"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	database.ConnectDb()
}

func main() {
	router.ClientRoutes()
}
