package pesanan

import (
    "time"
    "nasgorid_be/models/menu"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Pesanan struct {
    ID            primitive.ObjectID `bson:"_id,omitempty"`
	NamaPelanggan string    `json:"nama_pelanggan"`
	DaftarMenu    []menu.Menu    `json:"daftar_menu"`
	Jumlah        int       `json:"jumlah"`
	TotalHarga    float64   `json:"total_harga"`
	StatusPesanan string    `json:"status_pesanan"`
	WaktuPesanan  time.Time `json:"waktu_pesanan"`
}
