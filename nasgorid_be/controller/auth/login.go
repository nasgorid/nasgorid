package auth

import (
    "context"
    "encoding/json"
    "net/http"
    "nasgorid_be/models/pelanggan"
    "nasgorid_be/config"
    "go.mongodb.org/mongo-driver/bson"
    "golang.org/x/crypto/bcrypt"
)

// Login handles the login process
func Login(w http.ResponseWriter, r *http.Request) {
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    // Parse request body
    err := json.NewDecoder(r.Body).Decode(&credentials)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Cari pelanggan berdasarkan email
    collection := config.ConnectDB().Collection("pelanggan")
    var p pelanggan.Pelanggan
    err = collection.FindOne(context.TODO(), bson.M{"email": credentials.Email}).Decode(&p)
    if err != nil {
        http.Error(w, "Email atau password salah", http.StatusUnauthorized)
        return
    }

    // Bandingkan password
    err = bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(credentials.Password))
    if err != nil {
        http.Error(w, "Email atau password salah", http.StatusUnauthorized)
        return
    }

    // Jika login berhasil, kirim respons sukses
    json.NewEncoder(w).Encode(map[string]string{"message": "Login berhasil", "email": p.Email})
}
