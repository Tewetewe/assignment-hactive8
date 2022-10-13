package main

import (
	"assignment4/database"
	"assignment4/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
