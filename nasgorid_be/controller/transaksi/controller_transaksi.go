package transaksi

import (
    "context"
    "fmt"
    "log"
    "time"

    "nasgorid/models/transaksi" // Sesuaikan dengan path model transaksi kamu
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

// GetAllTransaksi retrieves all the transactions from the "transaksi" collection
func GetAllTransaksi(transaksiCollection *mongo.Collection) {
    var daftarTransaksi []transaksi.Transaksi

    // Query untuk mendapatkan semua data di collection "transaksi"
    cursor, err := transaksiCollection.Find(context.TODO(), bson.D{})
    if err != nil {
        log.Fatal("Error finding documents: ", err)
    }
    log.Println("Koneksi ke collection 'transaksi' berhasil.")
    defer cursor.Close(context.TODO())

    // Decode hasil cursor ke dalam slice daftarTransaksi
    for cursor.Next(context.TODO()) {
        var t transaksi.Transaksi
        if err = cursor.Decode(&t); err != nil {
            log.Fatal("Error decoding document: ", err)
        }
        daftarTransaksi = append(daftarTransaksi, t)
    }

    // Cek error setelah iterasi
    if err := cursor.Err(); err != nil {
        log.Fatal("Cursor error: ", err)
    }

    // Tampilkan semua transaksi yang ditemukan
    if len(daftarTransaksi) == 0 {
        fmt.Println("Tidak ada transaksi yang ditemukan.")
    } else {
        fmt.Println("Daftar Transaksi:")
        for _, t := range daftarTransaksi {
            fmt.Printf("ID Pesanan: %s, Total Bayar: %.2f, Metode: %s, Waktu: %s\n", t.IDPesanan.Hex(), t.TotalBayar, t.MetodePembayaran, t.WaktuTransaksi.Format(time.RFC3339))
        }
    }
}

// GetTransaksiByID retrieves a single transaction by its ID
func GetTransaksiByID(id primitive.ObjectID, transaksiCollection *mongo.Collection) (*transaksi.Transaksi, error) {
    var t transaksi.Transaksi
    err := transaksiCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&t)
    if err != nil {
        log.Printf("Error finding transaksi: %v", err)
        return nil, err
    }
    return &t, nil
}

// InsertTransaksi inserts a new transaction into the "transaksi" collection
func InsertTransaksi(t transaksi.Transaksi, transaksiCollection *mongo.Collection) error {
    t.WaktuTransaksi = time.Now() // Set waktu transaksi saat pembuatan

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := transaksiCollection.InsertOne(ctx, t)
    if err != nil {
        log.Printf("Error inserting transaksi: %v", err)
        return err
    }

    fmt.Println("Transaksi berhasil ditambahkan")
    return nil
}

// UpdateTransaksi updates an existing transaction by its ID
func UpdateTransaksi(id primitive.ObjectID, updateData bson.M, transaksiCollection *mongo.Collection) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Filter by ID and update fields
    _, err := transaksiCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": updateData})
    if err != nil {
        log.Printf("Error updating transaksi: %v", err)
        return err
    }

    fmt.Println("Transaksi berhasil diperbarui")
    return nil
}

// DeleteTransaksi deletes a transaction by its ID
func DeleteTransaksi(id primitive.ObjectID, transaksiCollection *mongo.Collection) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := transaksiCollection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        log.Printf("Error deleting transaksi: %v", err)
        return err
    }

    fmt.Println("Transaksi berhasil dihapus")
    return nil
}
