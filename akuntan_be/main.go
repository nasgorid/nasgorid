// main.go
package main

import (
	"log"
	"net/http"

	"akuntan/config"
	"akuntan/router"
)

// Middleware untuk menangani CORS
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Jika request method adalah OPTIONS, kirim status OK langsung
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

	// Pasang middleware CORS
    r.Use(enableCORS)

	// Jalankan server di port 8080
	log.Println("Server is running on port 8081...")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
