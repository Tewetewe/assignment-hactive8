package main

import (
	"assignment2/database"
	"assignment2/routers"
)

func main() {

	defer database.StartDB().Close()

	var PORT = ":8080"

	routers.StartServer().Run(PORT)

}
