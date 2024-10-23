// main.go
package main

import (
	"log"
	"net/http"

	"akuntan/config"
	"akuntan/router"
)

// Middleware untuk menangani CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling CORS for request:", r.Method, r.URL.Path)

		// Mengizinkan semua origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Mengizinkan method-method yang diizinkan
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// Mengizinkan headers tertentu
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Mengizinkan request OPTIONS untuk preflight CORS
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Inisialisasi koneksi ke MongoDB
	config.InitMongoDB()

	// Setup Router
	r := router.SetupRouter()

	// Terapkan middleware CORS ke semua rute
	r.Use(corsMiddleware)

	
	// Jalankan server di port 8080
	log.Println("Server is running on port 8081...")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
