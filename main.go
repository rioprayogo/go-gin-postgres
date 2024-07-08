package main

import (
	"log"

	"go-gin-postgres/routes"
	"go-gin-postgres/utils"
)

func main() {
	// Inisialisasi koneksi database
	utils.InitDB()

	// Setup router dan jalankan server
	r := routes.SetupRouter()
	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
