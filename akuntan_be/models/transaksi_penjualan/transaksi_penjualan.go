package transaksi_penjualan

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SalesTransaction adalah struct untuk transaksi penjualan
type SalesTransaction struct {
    ID            primitive.ObjectID    `bson:"_id,omitempty" json:"id,omitempty"`
    TransactionDate time.Time `bson:"transactionDate" json:"transactionDate"`
    CustomerName  string    `bson:"customer_name" json:"customer_name"`
    Products      []Product `bson:"products" json:"products"`
    TotalAmount   float64   `bson:"total_amount" json:"total_amount"`
    PaymentMethod string    `bson:"payment_method" json:"payment_method"`
    PaymentStatus string    `bson:"payment_status" json:"payment_status"`
}

// Product adalah struct untuk produk dalam transaksi penjualan
type Product struct {
    ID          primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
    Name        string  `bson:"name" json:"name"`
    Price       float64 `bson:"price" json:"price"`
    Category    string  `bson:"category" json:"category"`
    Description string  `bson:"description" json:"description"`
    Stock       int     `bson:"stock" json:"stock"`
}