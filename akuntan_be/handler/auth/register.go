// handler/user.go
package handler

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "akuntan/config"
    "akuntan/models/user"
    "golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var user user.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Hash password sebelum disimpan
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }
    user.Password = string(hashedPassword)

    // Insert user ke MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = config.UserCollection.InsertOne(ctx, user)
    if err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    // Kirim respon sukses
    w.WriteHeader(http.StatusCreated)
    response := map[string]string{"message": "User created successfully"}
    json.NewEncoder(w).Encode(response)
}
