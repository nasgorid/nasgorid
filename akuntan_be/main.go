// main.go
package main

import (
	"log"
	"net/http"

	"akuntan/config"
	"akuntan/router"
)

func main() {
	// Inisialisasi koneksi ke MongoDB
	config.InitMongoDB()

	// Setup Router
	r := router.SetupRouter()

	// Jalankan server di port 8080
	log.Println("Server is running on port 8081...")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
