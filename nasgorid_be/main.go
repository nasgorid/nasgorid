package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"nasgorid_be/config"          // Sesuaikan dengan package config kamu
	"nasgorid_be/controller/auth" // Tambahkan ini jika kamu ingin menggunakan fungsi register
	"nasgorid_be/models/menu" // Sesuaikan dengan package model/menu kamu

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Mengizinkan semua origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Mengizinkan metode yang diinginkan
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Jika request adalah OPTIONS, balas dengan status 200
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}

func main() {
    // Menghubungkan ke MongoDB menggunakan fungsi dari config
    db := config.ConnectDB()

    // Akses collection "menu" dari database
    menuCollection := db.Collection("menu")

    // Membuat router dengan mux
    r := mux.NewRouter()

    // Endpoint untuk menampilkan semua data menu
    r.HandleFunc("/menu", func(w http.ResponseWriter, r *http.Request) {
        showAllMenu(menuCollection, w, r)
    }).Methods("GET")

    // Endpoint untuk registrasi pelanggan
    r.HandleFunc("/register", auth.Register).Methods("POST")

    // Endpoint untuk registrasi pelanggan
    r.HandleFunc("/login", auth.Login).Methods("POST")

    // Jalankan server di port 8080
    log.Println("Server running at http://localhost:8081")
    log.Fatal(http.ListenAndServe(":8081", enableCors(r)))
}

// Fungsi untuk menampilkan semua data dari collection "menu"
func showAllMenu(menuCollection *mongo.Collection, w http.ResponseWriter, _ *http.Request) {
    var menus []menu.Menu

    // Query untuk mendapatkan semua data di collection "menu"
    cursor, err := menuCollection.Find(context.TODO(), bson.D{})
    if err != nil {
        log.Fatal("Error finding documents: ", err)
        http.Error(w, "Gagal mengambil data menu", http.StatusInternalServerError)
        return
    }
    log.Println("Koneksi ke collection 'menu' berhasil.")
    defer cursor.Close(context.TODO())

    // Decode hasil cursor ke dalam slice menus
    for cursor.Next(context.TODO()) {
        var menu menu.Menu
        if err = cursor.Decode(&menu); err != nil {
            log.Fatal("Error decoding document: ", err)
            http.Error(w, "Error decoding menu data", http.StatusInternalServerError)
            return
        }
        menus = append(menus, menu)
    }

    // Cek error setelah iterasi
    if err := cursor.Err(); err != nil {
        log.Fatal("Cursor error: ", err)
        http.Error(w, "Cursor error", http.StatusInternalServerError)
        return
    }

    // Tampilkan semua menu yang ditemukan
    if len(menus) == 0 {
        fmt.Fprintln(w, "Tidak ada menu yang ditemukan.")
    } else {
        fmt.Fprintln(w, "Daftar Menu:")
        for _, menu := range menus {
            fmt.Fprintf(w, "Nama: %s, Harga: %.2f, Tersedia: %t\n", menu.NamaMenu, menu.Harga, menu.Tersedia)
        }
    }
}


