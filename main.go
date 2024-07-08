package main

import (
	"log"

	"your-module/routes"
	"your-module/utils"
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
