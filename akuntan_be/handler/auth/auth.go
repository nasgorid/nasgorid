// handler/user.go
package auth

import (
	"context"
	"encoding/json"
	"net/http"
	// "strings"
	"time"

	"akuntan/config"
	"akuntan/models/user"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key") // Gantilah dengan secret key Anda

// Struct untuk JWT claims
type Claims struct {
    Email string `json:"email"`
    jwt.RegisteredClaims
}


func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var newUser user.User // Menggunakan tipe user.User yang diimpor
    if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Hash password sebelum disimpan
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }
    newUser.Password = string(hashedPassword)

    // Insert user ke MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = config.UserCollection.InsertOne(ctx, newUser)
    if err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    // Kirim respon sukses
    w.WriteHeader(http.StatusCreated)
    response := map[string]string{"message": "User created successfully"}
    json.NewEncoder(w).Encode(response)
}


func LoginUser(w http.ResponseWriter, r *http.Request) {
    var loginUser user.User // Menggunakan model User
    if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Cek apakah user dengan email tersebut ada di database
    var foundUser user.User
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    err := config.UserCollection.FindOne(ctx, bson.M{"email": loginUser.Email}).Decode(&foundUser)
    if err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    // Cek apakah password yang dimasukkan cocok dengan yang di-hash di database
    err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginUser.Password))
    if err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    // Kirim respon sukses
    response := map[string]string{"message": "Login successful"}
    json.NewEncoder(w).Encode(response)

    // // Buat JWT token
    // expirationTime := time.Now().Add(24 * time.Hour) // Token berlaku selama 24 jam
    // claims := &Claims{
    //     Email: foundUser.Email,
    //     RegisteredClaims: jwt.RegisteredClaims{
    //         ExpiresAt: jwt.NewNumericDate(expirationTime),
    //     },
    // }

    // token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    // tokenString, err := token.SignedString(jwtKey)
    // if err != nil {
    //     http.Error(w, "Failed to generate token", http.StatusInternalServerError)
    //     return
    // }

    // // Kirim token sebagai respon
    // w.Header().Set("Content-Type", "application/json")
    // json.NewEncoder(w).Encode(map[string]string{
    //     "token": tokenString,
    // })
}


// Definisikan tipe baru untuk context key
// type contextKey string

// // Middleware untuk memvalidasi token JWT
// func AuthMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         authHeader := r.Header.Get("Authorization")
//         if authHeader == "" {
//             http.Error(w, "Missing token", http.StatusUnauthorized)
//             return
//         }

//         parts := strings.Split(authHeader, " ")
//         if len(parts) != 2 || parts[0] != "Bearer" {
//             http.Error(w, "Invalid token format", http.StatusUnauthorized)
//             return
//         }

//         tokenString := parts[1]

//         // Validasi token JWT
//         claims := &Claims{}
//         token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
//             return jwtKey, nil
//         })

//         if err != nil || !token.Valid {
//             http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
//             return
//         }

//         // Gunakan contextKey sebagai tipe key untuk context
//         emailKey := contextKey("email")
//         ctx := context.WithValue(r.Context(), emailKey, claims.Email)
//         next.ServeHTTP(w, r.WithContext(ctx))
//     })
// }



