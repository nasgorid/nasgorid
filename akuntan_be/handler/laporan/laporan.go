package laporan

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"akuntan/config"
	"akuntan/models/laporan"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Handler untuk membuat laporan keuangan
func CreateFinancialReport(w http.ResponseWriter, r *http.Request) {
	var newReport laporan.Laporan

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&newReport); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse tanggal dari string ke time.Time
	startDate, err := time.Parse("2006-01-02", newReport.StartDate) // Mengambil dari objek
	if err != nil {
		http.Error(w, "Invalid start date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	endDate, err := time.Parse("2006-01-02", newReport.EndDate) // Mengambil dari objek
	if err != nil {
		http.Error(w, "Invalid end date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	// Set waktu pembuatan
	newReport.StartDateTime = startDate // Pastikan field ini ada di model
	newReport.EndDateTime = endDate     // Pastikan field ini ada di model
	newReport.CreatedAt = time.Now()

	// Simpan ke MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newReport.ID = primitive.NewObjectID()
	_, err = config.ReportCollection.InsertOne(ctx, newReport)
	if err != nil {
		http.Error(w, "Failed to create financial report", http.StatusInternalServerError)
		return
	}

	// Kirim respon sukses
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Financial report created successfully"})
}

// Fungsi untuk mendapatkan laporan keuangan berdasarkan ID
func GetFinancialReportByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid report ID", http.StatusBadRequest)
		return
	}

	var report laporan.Laporan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = config.ReportCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&report)
	if err != nil {
		http.Error(w, "Financial report not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

// Handler untuk mendapatkan semua laporan keuangan
func GetFinancialReports(w http.ResponseWriter, r *http.Request) {
	var reports []laporan.Laporan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := config.ReportCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch financial reports", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var rep laporan.Laporan
		if err := cursor.Decode(&rep); err != nil {
			http.Error(w, "Error decoding financial report", http.StatusInternalServerError)
			return
		}
		reports = append(reports, rep)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}

// Fungsi untuk menghapus laporan keuangan berdasarkan ID
func DeleteFinancialReport(w http.ResponseWriter, r *http.Request) {
	// Ambil parameter ID dari URL menggunakan gorilla/mux
	vars := mux.Vars(r)
	id := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid report ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = config.ReportCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, "Failed to delete financial report", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Financial report deleted successfully"})
}
