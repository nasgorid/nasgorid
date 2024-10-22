package produk

// Product adalah struct untuk produk
type Product struct {
    ID          string  `bson:"_id,omitempty" json:"id,omitempty"`
    Name        string  `bson:"name" json:"name"`
    Price       float64 `bson:"price" json:"price"`
    Category    string  `bson:"category" json:"category"`
    Description string  `bson:"description" json:"description"`
    Stock       int     `bson:"stock" json:"stock"`
    CreatedAt   int64   `bson:"createdAt" json:"createdAt"`
    UpdatedAt   int64   `bson:"updatedAt" json:"updatedAt"`
}