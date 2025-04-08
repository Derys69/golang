package main

import (
	"golang/database"
	"golang/routes"
)

func main() {
	database.InitDB()         // koneksi ke database
	r := routes.SetupRouter() // setup router
	r.Run(":8080")            // koneksi server pada port 8080
}
