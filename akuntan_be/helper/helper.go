package helper

import (
	"time"
	"log"
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


