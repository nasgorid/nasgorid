package pesanan

import (
	"context"
	"fmt"
	"log"
	"time"

	"nasgorid_be/models/pesanan" // Sesuaikan dengan path model pesanan kamu

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllPesanan retrieves all the orders from the "pesanan" collection
func GetAllPesanan(pesananCollection *mongo.Collection) {
	var daftarPesanan []pesanan.Pesanan

	// Query untuk mendapatkan semua data di collection "pesanan"
	cursor, err := pesananCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal("Error finding documents: ", err)
	}
	log.Println("Koneksi ke collection 'pesanan' berhasil.")
	defer cursor.Close(context.TODO())

	// Decode hasil cursor ke dalam slice daftarPesanan
	for cursor.Next(context.TODO()) {
		var p pesanan.Pesanan
		if err = cursor.Decode(&p); err != nil {
			log.Fatal("Error decoding document: ", err)
		}
		daftarPesanan = append(daftarPesanan, p)
	}

	// Cek error setelah iterasi
	if err := cursor.Err(); err != nil {
		log.Fatal("Cursor error: ", err)
	}

	// Tampilkan semua pesanan yang ditemukan
	if len(daftarPesanan) == 0 {
		fmt.Println("Tidak ada pesanan yang ditemukan.")
	} else {
		fmt.Println("Daftar Pesanan:")
		for _, p := range daftarPesanan {
			fmt.Printf("Nama Pelanggan: %s, Total Harga: %.2f, Status: %s\n", p.NamaPelanggan, p.TotalHarga, p.StatusPesanan)
		}
	}
}

// GetPesananByID retrieves a single order by its ID
func GetPesananByID(id primitive.ObjectID, pesananCollection *mongo.Collection) (*pesanan.Pesanan, error) {
	var p pesanan.Pesanan
	err := pesananCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&p)
	if err != nil {
		log.Printf("Error finding pesanan: %v", err)
		return nil, err
	}
	return &p, nil
}

// InsertPesanan inserts a new order into the "pesanan" collection
func InsertPesanan(p pesanan.Pesanan, pesananCollection *mongo.Collection) error {
	p.WaktuPesanan = time.Now() // Set waktu pesanan pada saat pembuatan

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := pesananCollection.InsertOne(ctx, p)
	if err != nil {
		log.Printf("Error inserting pesanan: %v", err)
		return err
	}

	fmt.Println("Pesanan berhasil ditambahkan")
	return nil
}

// UpdatePesanan updates an existing order by its ID
func UpdatePesanan(id primitive.ObjectID, updateData bson.M, pesananCollection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filter by ID and update fields
	_, err := pesananCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": updateData})
	if err != nil {
		log.Printf("Error updating pesanan: %v", err)
		return err
	}

	fmt.Println("Pesanan berhasil diperbarui")
	return nil
}

// DeletePesanan deletes an order by its ID
func DeletePesanan(id primitive.ObjectID, pesananCollection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := pesananCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Printf("Error deleting pesanan: %v", err)
		return err
	}

	fmt.Println("Pesanan berhasil dihapus")
	return nil
}
