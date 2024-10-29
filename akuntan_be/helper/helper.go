// helper/helper.go
package helper

import (
	"log"
	"strconv"
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

// FormatCurrency mengonversi angka menjadi format mata uang.
func FormatCurrency(amount float64) string {
	return strconv.FormatFloat(amount, 'f', 2, 64) + " IDR" // Ganti "IDR" dengan simbol mata uang yang diinginkan.
}

// GetCurrentTime mengembalikan waktu saat ini.
func GetCurrentTime() time.Time {
	return time.Now()
}

// Fungsi statistik
// CalculateTotal menghitung total dari sebuah slice of float64
func CalculateTotal(numbers []float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total
}

// CalculateAverage menghitung rata-rata dari sebuah slice of float64
func CalculateAverage(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	total := CalculateTotal(numbers)
	return total / float64(len(numbers))
}
