package transaksi_pengeluaran

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "akuntan/config"
    "akuntan/models/transaksi_pengeluaran"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateExpenseTransaction adalah handler untuk membuat transaksi pengeluaran
func CreateExpenseTransaction(w http.ResponseWriter, r *http.Request) {
    var newExpense transaksi_pengeluaran.ExpenseTransaction
    if err := json.NewDecoder(r.Body).Decode(&newExpense); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    newExpense.CreatedAt = time.Now()
    newExpense.UpdatedAt = time.Now()

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := config.ExpenseTransactionCollection.InsertOne(ctx, newExpense)
    if err != nil {
        http.Error(w, "Failed to create expense transaction", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Expense transaction created successfully"})
}

// GetExpenses mengembalikan semua transaksi pengeluaran
func GetExpenses(w http.ResponseWriter, r *http.Request) {
    var expenses []transaksi_pengeluaran.ExpenseTransaction

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := config.ExpenseTransactionCollection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, "Failed to fetch expense transactions", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var exp transaksi_pengeluaran.ExpenseTransaction
        if err := cursor.Decode(&exp); err != nil {
            http.Error(w, "Error decoding expense transaction", http.StatusInternalServerError)
            return
        }
        expenses = append(expenses, exp)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(expenses)
}

// GetExpenseByID mengambil transaksi pengeluaran berdasarkan ID
func GetExpenseByID(w http.ResponseWriter, r *http.Request) {
    // id := r.URL.Query().Get("id")

    vars := mux.Vars(r)
	id := vars["id"]

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, "Invalid expense transaction ID", http.StatusBadRequest)
        return
    }

    var expense transaksi_pengeluaran.ExpenseTransaction
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    err = config.ExpenseTransactionCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&expense)
    if err != nil {
        http.Error(w, "Expense transaction not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(expense)
}

// UpdateExpense mengupdate transaksi pengeluaran berdasarkan ID
func UpdateExpense(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, "Invalid expense transaction ID", http.StatusBadRequest)
        return
    }

    var updatedExpense transaksi_pengeluaran.ExpenseTransaction
    if err := json.NewDecoder(r.Body).Decode(&updatedExpense); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    update := bson.M{
        "$set": bson.M{
            "expense_name":  updatedExpense.ExpenseName,
            "amount":        updatedExpense.Amount,
            "category":      updatedExpense.Category,
            "payment_method": updatedExpense.PaymentMethod,
            "expense_date":  updatedExpense.ExpenseDate,
            "notes":         updatedExpense.Notes,
            "updated_at":    time.Now(),
        },
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = config.ExpenseTransactionCollection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
    if err != nil {
        http.Error(w, "Failed to update expense transaction", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Expense transaction updated successfully"})
}

// DeleteExpense menghapus transaksi pengeluaran berdasarkan ID
func DeleteExpense(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, "Invalid expense transaction ID", http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = config.ExpenseTransactionCollection.DeleteOne(ctx, bson.M{"_id": objectID})
    if err != nil {
        http.Error(w, "Failed to delete expense transaction", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Expense transaction deleted successfully"})
}
