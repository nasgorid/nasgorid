package transaksi_pengeluaran

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ExpenseTransaction adalah struct untuk transaksi pengeluaran
type ExpenseTransaction struct {
    ID            primitive.ObjectID    `bson:"_id,omitempty" json:"id,omitempty"`
    ExpenseName  string    `bson:"expense_name" json:"expense_name"`   // Nama pengeluaran (misalnya: sewa, gaji, dll.)
    Amount        float64   `bson:"amount" json:"amount"`               // Jumlah uang yang dikeluarkan
    Category      string    `bson:"category" json:"category"`           // Kategori pengeluaran (misalnya: operasional, marketing, dll.)
    PaymentMethod string   `bson:"payment_method" json:"payment_method"` // Metode pembayaran (misalnya: transfer bank, tunai)
    ExpenseDate  time.Time `bson:"expense_date" json:"expense_date"`   // Tanggal pengeluaran
    Notes         string    `bson:"notes" json:"notes,omitempty"`       // Catatan tambahan (opsional)
    CreatedAt    time.Time `bson:"created_at" json:"created_at"`       // Waktu transaksi dibuat
    UpdatedAt    time.Time `bson:"updated_at" json:"updated_at"`       // Waktu transaksi terakhir diperbarui
}
