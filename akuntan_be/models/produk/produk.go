package produk

import "time"

// Product adalah struct untuk produk
type Product struct {
    ID          string  `bson:"_id,omitempty" json:"id,omitempty"`
    Name        string  `bson:"name" json:"name"`
    Price       float64 `bson:"price" json:"price"`
    Category    string  `bson:"category" json:"category"`
    Description string  `bson:"description" json:"description"`
    Stock       int     `bson:"stock" json:"stock"`
    CreatedAt   time.Time   `bson:"createdAt" json:"createdAt"`
    UpdatedAt   time.Time   `bson:"updatedAt" json:"updatedAt"`
}