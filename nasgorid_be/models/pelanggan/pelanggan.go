package pelanggan

import "go.mongodb.org/mongo-driver/bson/primitive"

// Pelanggan represents a customer or user in the system
type Pelanggan struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Nama      string             `bson:"nama" json:"nama"`
    Email     string             `bson:"email" json:"email"`
    NoTelp    string             `bson:"no_telp" json:"no_telp"`
    Alamat    string             `bson:"alamat" json:"alamat"`
    Password  string             `bson:"password" json:"password"`
    CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
    UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
}