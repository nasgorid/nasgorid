package transaksi

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"

)

type Transaksi struct {
    ID             primitive.ObjectID `bson:"_id,omitempty"`
    IDPesanan      primitive.ObjectID `bson:"id_pesanan"`
	TotalBayar       float64   `json:"total_bayar"`
	MetodePembayaran string    `json:"metode_pembayaran"`
	WaktuTransaksi   time.Time `json:"waktu_transaksi"`
}
