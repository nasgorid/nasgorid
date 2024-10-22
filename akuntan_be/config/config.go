// database/connection.go
package config

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var UserCollection *mongo.Collection
var ProductCollection *mongo.Collection

func InitMongoDB() {
    // Gunakan MongoDB URI yang diberikan
    uri := "mongodb+srv://karamissuu:karamissu1@cluster0.lyovb.mongodb.net/?retryWrites=true&w=majority"

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal(err)
    }

    log.Println("MongoDB connected")

    Client = client
    UserCollection = Client.Database("akuntan").Collection("user")
}
