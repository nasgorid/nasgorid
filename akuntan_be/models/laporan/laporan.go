package laporan

type Laporan struct {
    ID        string  `bson:"_id,omitempty" json:"id,omitempty"`
    StartDate int64   `bson:"startDate" json:"startDate"`
    EndDate   int64   `bson:"endDate" json:"endDate"`
    Income    float64 `bson:"income" json:"income"`
    Expenses  float64 `bson:"expenses" json:"expenses"`
    Profit    float64 `bson:"profit" json:"profit"`
    CreatedAt int64   `bson:"createdAt" json:"createdAt"`
}
