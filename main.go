package main

import (
	"golang/database"
	"golang/routes"
)

func main() {
	database.InitDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
