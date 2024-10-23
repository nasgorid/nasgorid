package pelanggan

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"akuntan/config"
	"akuntan/models/pelanggan"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateCustomer handles creating a new customer
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var newCustomer pelanggan.Customer
	if err := json.NewDecoder(r.Body).Decode(&newCustomer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newCustomer.CreatedAt = time.Now()
	newCustomer.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.CustomerCollection.InsertOne(ctx, newCustomer)
	if err != nil {
		http.Error(w, "Failed to create customer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer created successfully"})
}

// GetCustomers handles retrieving all customers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []pelanggan.Customer

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := config.CustomerCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch customers", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var cust pelanggan.Customer
		if err := cursor.Decode(&cust); err != nil {
			http.Error(w, "Error decoding customer", http.StatusInternalServerError)
			return
		}
		customers = append(customers, cust)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

// GetCustomerByID handles retrieving a customer by ID
func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	var customer pelanggan.Customer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = config.CustomerCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&customer)
	if err != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// UpdateCustomer handles updating a customer by ID
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	// Ambil parameter ID dari URL menggunakan gorilla/mux
	vars := mux.Vars(r)
	id := vars["id"]
	
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	var updatedCustomer pelanggan.Customer
	if err := json.NewDecoder(r.Body).Decode(&updatedCustomer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":      updatedCustomer.Name,
			"email":     updatedCustomer.Email,
			"phone":     updatedCustomer.Phone,
			"address":   updatedCustomer.Address,
			"updatedAt": time.Now(),
		},
	}

	_, err = config.CustomerCollection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		http.Error(w, "Failed to update customer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer updated successfully"})
}

// DeleteCustomer handles deleting a customer by ID
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	// Ambil parameter ID dari URL menggunakan gorilla/mux
	vars := mux.Vars(r)
	id := vars["id"]
	
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = config.CustomerCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, "Failed to delete customer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer deleted successfully"})
}
