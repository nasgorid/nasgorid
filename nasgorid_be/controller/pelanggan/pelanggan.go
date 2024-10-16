package pelanggan

import (
	"context"
	"encoding/json"
	"nasgorid_be/config"
	"nasgorid_be/models/pelanggan"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePelanggan handles the creation of a new pelanggan
func CreatePelanggan(w http.ResponseWriter, r *http.Request) {
    var pelanggan pelanggan.Pelanggan
    json.NewDecoder(r.Body).Decode(&pelanggan)

    pelanggan.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
    pelanggan.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

    collection := config.ConnectDB().Collection("pelanggan")
    _, err := collection.InsertOne(context.TODO(), pelanggan)
    if err != nil {
        http.Error(w, "Error creating pelanggan", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(pelanggan)
}