// handler/auth/login.go
package handler

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "akuntan/config"
    "akuntan/models/user"
    "go.mongodb.org/mongo-driver/bson"
    "golang.org/x/crypto/bcrypt"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
    var loginData struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Cari user berdasarkan email
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var user user.User
    err := config.UserCollection.FindOne(ctx, bson.M{"email": loginData.Email}).Decode(&user)
    if err != nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }

    // Cek password
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
    if err != nil {
        http.Error(w, "Invalid password", http.StatusUnauthorized)
        return
    }

    // Kirim respon sukses
    response := map[string]string{"message": "Login successful"}
    json.NewEncoder(w).Encode(response)
}
