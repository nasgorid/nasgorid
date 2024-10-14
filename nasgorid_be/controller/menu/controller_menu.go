package menu

import (
    "context"
    "fmt"
    "log"
	 // Sesuaikan dengan package config kamu
    "nasgorid/models/menu" // Sesuaikan dengan package model/menu kamu
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

func GetAllMenu(menuCollection *mongo.Collection) {
    var menus []menu.Menu

    // Query untuk mendapatkan semua data di collection "menu"
    cursor, err := menuCollection.Find(context.TODO(), bson.D{})
    if err != nil {
        log.Fatal("Error finding documents: ", err)
    }
    log.Println("Koneksi ke collection 'menu' berhasil.")
    defer cursor.Close(context.TODO())

    // Decode hasil cursor ke dalam slice menus
    for cursor.Next(context.TODO()) {
        var menu menu.Menu
        if err = cursor.Decode(&menu); err != nil {
            log.Fatal("Error decoding document: ", err)
        }
        menus = append(menus, menu)
    }

    // Cek error setelah iterasi
    if err := cursor.Err(); err != nil {
        log.Fatal("Cursor error: ", err)
    }

    // Tampilkan semua menu yang ditemukan
    if len(menus) == 0 {
        fmt.Println("Tidak ada menu yang ditemukan.")
    } else {
        fmt.Println("Daftar Menu:")
        for _, menu := range menus {
            fmt.Printf("Nama: %s, Harga: %.2f, Tersedia: %t\n", menu.NamaMenu, menu.Harga, menu.Tersedia)
        }
    }
}
