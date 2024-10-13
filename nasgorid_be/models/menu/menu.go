package menu

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    NamaMenu  string             `bson:"nama_menu"`
    Harga     float64            `bson:"harga"`
    Deskripsi string             `bson:"deskripsi,omitempty"`
    Gambar    string             `bson:"gambar,omitempty"`
    Tersedia  bool               `bson:"tersedia"`
}
