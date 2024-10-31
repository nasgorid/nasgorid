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

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// var jwtKey = []byte("your_secret_key") // Gantilah dengan secret key Anda

// // Struct untuk JWT claims
// type Claims struct {
//     Email string `json:"email"`
//     jwt.RegisteredClaims
// }


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

// // Fungsi untuk mendapatkan daftar produk
// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	var users []user.User

// 	// Ambil data dari MongoDB
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	cursor, err := config.UserCollection.Find(ctx, bson.M{})
// 	if err != nil {
// 		http.Error(w, "Failed to fetch User", http.StatusInternalServerError)
// 		return
// 	}
// 	defer cursor.Close(ctx)

// 	for cursor.Next(ctx) {
// 		var user user.User
// 		if err := cursor.Decode(&user); err != nil {
// 			http.Error(w, "Error decoding user", http.StatusInternalServerError)
// 			return
// 		}
// 		users = append(users, user)
// 	}

// 	// Kirim data user sebagai respon
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(users)
// }

// Fungsi untuk mendapatkan detail produk berdasarkan ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
    // Ambil parameter ID dari URL menggunakan mux.Vars
    vars := mux.Vars(r)
    id := vars["id"]

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Ambil produk dari MongoDB
    var user user.User
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    err = config.UserCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Kirim produk sebagai respon
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}


func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Ambil parameter ID dari URL menggunakan gorilla/mux
	vars := mux.Vars(r)
	id := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	// Decode data produk yang akan diupdate
	var updatedUser user.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update produk di MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":        updatedUser.Name,
			"email":       updatedUser.Email,
			"password":       updatedUser.Password,
			"umkmName":    updatedUser.UMKMName,
			"updatedAt":   time.Now(),
		},
	}

	_, err = config.UserCollection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		http.Error(w, "Failed to update User", http.StatusInternalServerError)
		return
	}

	// Kirim respon sukses
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

