package helper

import (
    "errors"
    "fmt"
    "log"
    "os"
)

// ValidateMenu validates if the menu data is complete
func ValidateMenu(name string, price float64, available bool) error {
    if name == "" {
        return errors.New("nama menu tidak boleh kosong")
    }
    if price <= 0 {
        return errors.New("harga menu harus lebih dari nol")
    }
    return nil
}

// FormatPrice formats the price into IDR currency
func FormatPrice(price float64) string {
    return fmt.Sprintf("Rp %.2f", price)
}

// LogError logs errors to a file or console
func LogError(err error) {
    if err != nil {
        // Tulis ke file log (opsional)
        logFile, _ := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        logger := log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
        logger.Println(err.Error())

        // Atau langsung tampilkan di console
        log.Println("ERROR:", err.Error())
    }
}
