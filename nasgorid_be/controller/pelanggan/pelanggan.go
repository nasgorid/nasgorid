package pelanggan

import (
	"context"
	"encoding/json"
	"nasgorid_be/config"
	"nasgorid_be/models/pelanggan"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
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

// GetPelangganByID fetches a single pelanggan by ID
func GetPelangganByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    var pelanggan pelanggan.Pelanggan
    collection := config.ConnectDB().Collection("pelanggan")

    err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&pelanggan)
    if err != nil {
        http.Error(w, "Pelanggan not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(pelanggan)
}

// UpdatePelanggan handles updating pelanggan data by ID
func UpdatePelanggan(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    var updatedData bson.M
    json.NewDecoder(r.Body).Decode(&updatedData)

    updatedData["updated_at"] = primitive.NewDateTimeFromTime(time.Now())

    collection := config.ConnectDB().Collection("pelanggan")
    _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": updatedData})
    if err != nil {
        http.Error(w, "Error updating pelanggan", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode("Pelanggan updated successfully")
}
