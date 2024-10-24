package laporan

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Laporan struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    StartDate  string             `bson:"startDate" json:"startDate"`   // Menggunakan string
    EndDate    string             `bson:"endDate" json:"endDate"`       // Menggunakan string
    StartDateTime time.Time       `bson:"startDateTime" json:"startDateTime"` // Menyimpan waktu yang sudah diparse
    EndDateTime   time.Time       `bson:"endDateTime" json:"endDateTime"`
	Income    float64   `bson:"income" json:"income"`
	Expenses  float64   `bson:"expenses" json:"expenses"`
	Profit    float64   `bson:"profit" json:"profit"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
