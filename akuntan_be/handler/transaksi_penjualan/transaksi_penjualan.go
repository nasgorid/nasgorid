package transaksi_penjualan

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"akuntan/config"
	"akuntan/models/transaksi_penjualan"

	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"encoding/csv"
	"fmt"
)

// Fungsi untuk menambahkan transaksi penjualan baru
func CreateSalesTransaction(w http.ResponseWriter, r *http.Request) {
	var newTransaction transaksi_penjualan.SalesTransaction
	if err := json.NewDecoder(r.Body).Decode(&newTransaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set waktu transaksi
	newTransaction.TransactionDate = time.Now()

	// Insert transaksi ke MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.SalesTransactionCollection.InsertOne(ctx, newTransaction)
	if err != nil {
		http.Error(w, "Failed to create sales transaction", http.StatusInternalServerError)
		return
	}

	// Kirim respon sukses
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sales transaction created successfully"})
}

// Fungsi untuk mendapatkan semua transaksi penjualan
func GetSalesTransactions(w http.ResponseWriter, r *http.Request) {
	var transactions []transaksi_penjualan.SalesTransaction

	// Ambil data dari MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := config.SalesTransactionCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch sales transactions", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var transaction transaksi_penjualan.SalesTransaction
		if err := cursor.Decode(&transaction); err != nil {
			http.Error(w, "Error decoding sales transaction", http.StatusInternalServerError)
			return
		}
		transactions = append(transactions, transaction)
	}

	// Kirim data transaksi sebagai respon
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

// Fungsi untuk mendapatkan transaksi penjualan berdasarkan ID
func GetSalesTransactionByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid sales transaction ID", http.StatusBadRequest)
		return
	}

	// Ambil transaksi dari MongoDB
	var transaction transaksi_penjualan.SalesTransaction
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = config.SalesTransactionCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&transaction)
	if err != nil {
		http.Error(w, "Sales transaction not found", http.StatusNotFound)
		return
	}

	// Kirim transaksi sebagai respon
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

// Fungsi untuk mengupdate transaksi penjualan berdasarkan ID
func UpdateSalesTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid sales transaction ID", http.StatusBadRequest)
		return
	}

	// Decode data transaksi yang akan diupdate
	var updatedTransaction transaksi_penjualan.SalesTransaction
	if err := json.NewDecoder(r.Body).Decode(&updatedTransaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update transaksi di MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"customerName":   updatedTransaction.CustomerName,
			"products":       updatedTransaction.Products,
			"totalAmount":    updatedTransaction.TotalAmount,
			"paymentMethod":  updatedTransaction.PaymentMethod,
			"paymentStatus":  updatedTransaction.PaymentStatus,
			"transactionDate": time.Now(),
		},
	}

	_, err = config.SalesTransactionCollection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		http.Error(w, "Failed to update sales transaction", http.StatusInternalServerError)
		return
	}

	// Kirim respon sukses
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sales transaction updated successfully"})
}

// Fungsi untuk menghapus transaksi penjualan berdasarkan ID
func DeleteSalesTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid sales transaction ID", http.StatusBadRequest)
		return
	}

	// Hapus transaksi dari MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = config.SalesTransactionCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, "Failed to delete sales transaction", http.StatusInternalServerError)
		return
	}

	// Kirim respon sukses
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sales transaction deleted successfully"})
}


// ExportSalesTransactionsToCSV mengexport semua transaksi penjualan ke CSV
func ExportSalesTransactionsCSV(w http.ResponseWriter, r *http.Request) {
    // Ambil semua data transaksi dari MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := config.SalesTransactionCollection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, "Failed to fetch sales transactions", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    // Buat file CSV
    file, err := os.Create("sales_transactions.csv")
    if err != nil {
        http.Error(w, "Failed to create CSV file", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Tulis header CSV
    header := []string{"ID", "Transaction Date", "Customer Name", "Total Amount", "Payment Method", "Payment Status"}
    if err := writer.Write(header); err != nil {
        http.Error(w, "Failed to write header to CSV", http.StatusInternalServerError)
        return
    }

    // Tulis data transaksi ke CSV
    for cursor.Next(ctx) {
        var transaction transaksi_penjualan.SalesTransaction
        if err := cursor.Decode(&transaction); err != nil {
            http.Error(w, "Error decoding sales transaction", http.StatusInternalServerError)
            return
        }

        record := []string{
            transaction.ID.Hex(), // Konversi ObjectID ke string
            transaction.TransactionDate.Format(time.RFC3339),
            transaction.CustomerName,
            fmt.Sprintf("%.2f", transaction.TotalAmount), // Format total amount
            transaction.PaymentMethod,
            transaction.PaymentStatus,
        }

        if err := writer.Write(record); err != nil {
            http.Error(w, "Failed to write transaction to CSV", http.StatusInternalServerError)
            return
        }
    }

    // Kirim respon sukses
    w.Header().Set("Content-Type", "text/csv")
    w.Header().Set("Content-Disposition", "attachment;filename=sales_transactions.csv")
    http.ServeFile(w, r, "sales_transactions.csv")
}