package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// ConnectDB initializes and returns a MongoDB client
func ConnectDB() *mongo.Database {
	var err error
	// Menggunakan URI yang telah diperbarui
	clientOptions := options.Client().ApplyURI("mongodb+srv://karamissuu:karamissu1@cluster0.lyovb.mongodb.net/nasgorid?retryWrites=true&w=majority")
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Menguji koneksi
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Berhasil terhubung ke database MongoDB")

	// Mengembalikan database yang digunakan
	return client.Database("nasgorid") // Pastikan nama database sesuai
}
