package pesanan

import (
	"context"
	"encoding/json"

	"net/http"
	"time"

	"nasgorid_be/models/pesanan" // Sesuaikan dengan path model pesanan kamu
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllPesanan retrieves all the orders from the "pesanan" collection
func GetAllPesanan(w http.ResponseWriter, r *http.Request, pesananCollection *mongo.Collection) {
	var daftarPesanan []pesanan.Pesanan

	// Query untuk mendapatkan semua data di collection "pesanan"
	cursor, err := pesananCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		http.Error(w, "Error finding documents", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	// Decode hasil cursor ke dalam slice daftarPesanan
	for cursor.Next(context.TODO()) {
		var p pesanan.Pesanan
		if err = cursor.Decode(&p); err != nil {
			http.Error(w, "Error decoding document", http.StatusInternalServerError)
			return
		}
		daftarPesanan = append(daftarPesanan, p)
	}

	// Cek error setelah iterasi
	if err := cursor.Err(); err != nil {
		http.Error(w, "Cursor error", http.StatusInternalServerError)
		return
	}

	// Encode hasil pesanan ke dalam bentuk JSON
	w.Header().Set("Content-Type", "application/json")
	if len(daftarPesanan) == 0 {
		json.NewEncoder(w).Encode([]pesanan.Pesanan{})
	} else {
		json.NewEncoder(w).Encode(daftarPesanan)
	}
}

// GetPesananByID retrieves a single order by its ID
func GetPesananByID(w http.ResponseWriter, r *http.Request, pesananCollection *mongo.Collection) {
	params := r.URL.Query()
	id, err := primitive.ObjectIDFromHex(params.Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var p pesanan.Pesanan
	err = pesananCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&p)
	if err != nil {
		http.Error(w, "Pesanan not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// InsertPesanan inserts a new order into the "pesanan" collection
func InsertPesanan(w http.ResponseWriter, r *http.Request, pesananCollection *mongo.Collection) {
	var p pesanan.Pesanan

	// Decode request body ke dalam struct pesanan
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	p.WaktuPesanan = time.Now() // Set waktu pesanan pada saat pembuatan

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := pesananCollection.InsertOne(ctx, p)
	if err != nil {
		http.Error(w, "Error inserting pesanan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Pesanan berhasil ditambahkan")
}

// UpdatePesanan updates an existing order by its ID
func UpdatePesanan(w http.ResponseWriter, r *http.Request, pesananCollection *mongo.Collection) {
	params := r.URL.Query()
	id, err := primitive.ObjectIDFromHex(params.Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedData bson.M
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = pesananCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": updatedData})
	if err != nil {
		http.Error(w, "Error updating pesanan", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Pesanan berhasil diperbarui")
}

// DeletePesanan deletes an order by its ID
func DeletePesanan(w http.ResponseWriter, r *http.Request, pesananCollection *mongo.Collection) {
	params := r.URL.Query()
	id, err := primitive.ObjectIDFromHex(params.Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = pesananCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		http.Error(w, "Error deleting pesanan", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Pesanan berhasil dihapus")
}
