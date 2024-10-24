package laporan

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "akuntan/config"
    "akuntan/models/laporan"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// Fungsi untuk membuat laporan keuangan baru
func CreateFinancialReport(w http.ResponseWriter, r *http.Request) {
    var newReport laporan.Laporan
    if err := json.NewDecoder(r.Body).Decode(&newReport); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    newReport.CreatedAt = time.Now().Unix()

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := config.ReportCollection.InsertOne(ctx, newReport)
    if err != nil {
        http.Error(w, "Failed to create financial report", http.StatusInternalServerError)
        return
    }

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

// Fungsi untuk mendapatkan semua laporan keuangan
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
    id := r.URL.Query().Get("id")
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
