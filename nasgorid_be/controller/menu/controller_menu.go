package menu

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	// Sesuaikan dengan package config kamu
	"nasgorid_be/config"
	"nasgorid_be/models/menu" // Sesuaikan dengan package model/menu kamu
	"nasgorid_be/models/pelanggan"

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

// InsertMenu inserts a new menu item into the "menu" collection
func InsertMenu(menu menu.Menu, db *mongo.Database) error {
    menuCollection := db.Collection("menu")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := menuCollection.InsertOne(ctx, menu)
    if err != nil {
        log.Printf("Error inserting menu: %v", err)
        return err
    }

    fmt.Println("Menu berhasil ditambahkan")
    return nil
}

// UpdateMenu updates a menu item based on its ID
func UpdateMenu(id string, updatedData bson.M, db *mongo.Database) error {
    menuCollection := db.Collection("menu")

    filter := bson.M{"_id": id}
    update := bson.M{"$set": updatedData}

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := menuCollection.UpdateOne(ctx, filter, update)
    if err != nil {
        log.Printf("Error updating menu: %v", err)
        return err
    }

    fmt.Println("Menu berhasil diperbarui")
    return nil
}

// DeleteMenu deletes a menu item based on its ID
func DeleteMenu(id string, db *mongo.Database) error {
    menuCollection := db.Collection("menu")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := menuCollection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        log.Printf("Error deleting menu: %v", err)
        return err
    }

    fmt.Println("Menu berhasil dihapus")
    return nil
}

// GetAllPelanggan fetches all pelanggan from the database
func GetAllPelanggan(w http.ResponseWriter, r *http.Request) {
    var pelanggans []pelanggan.Pelanggan
    collection := config.ConnectDB().Collection("pelanggan")

    cursor, err := collection.Find(context.TODO(), bson.D{})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var pelanggan pelanggan.Pelanggan
        cursor.Decode(&pelanggan)
        pelanggans = append(pelanggans, pelanggan)
    }

    json.NewEncoder(w).Encode(pelanggans)
}