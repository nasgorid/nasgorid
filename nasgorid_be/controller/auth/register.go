package auth

import (
    "context"
    "encoding/json"
    "net/http"
    "nasgorid/models/pelanggan"
    "nasgorid/config"
    "go.mongodb.org/mongo-driver/bson"
    "golang.org/x/crypto/bcrypt"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

// Register handles the registration of a new user
func Register(w http.ResponseWriter, r *http.Request) {
    var p pelanggan.Pelanggan // pastikan penggunaan struct Pelanggan benar
    json.NewDecoder(r.Body).Decode(&p)

    // Validasi apakah email sudah terdaftar
    collection := config.ConnectDB().Collection("pelanggan")
    var existingPelanggan pelanggan.Pelanggan
    err := collection.FindOne(context.TODO(), bson.M{"email": p.Email}).Decode(&existingPelanggan)
    if err == nil {
        http.Error(w, "Email sudah terdaftar", http.StatusBadRequest)
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Gagal mengenkripsi password", http.StatusInternalServerError)
        return
    }
    p.Password = string(hashedPassword)

    // Set tanggal pembuatan dan pembaruan
    p.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
    p.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

    // Masukkan pelanggan baru ke database
    _, err = collection.InsertOne(context.TODO(), p)
    if err != nil {
        http.Error(w, "Gagal mendaftarkan pelanggan", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(p)
}

