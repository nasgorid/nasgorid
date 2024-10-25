package helper

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mengonversi string ke time.Time dengan format "2006-01-02"
func StringToDate(dateStr string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Printf("Error parsing date: %v", err)
		return time.Time{}, err
	}
	return date, nil
}

// Mengonversi time.Time ke string
func DateToString(date time.Time) string {
	return date.Format("2006-01-02")
}

// Validasi ID dan konversi ke ObjectID
func ValidateID(id string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error converting ID: %v", err)
		return primitive.NilObjectID, err
	}
	return objectID, nil
}
