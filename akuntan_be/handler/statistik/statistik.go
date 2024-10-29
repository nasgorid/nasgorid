// handler/statistics.go
package statistik

import (
	"encoding/json"
	"net/http"
	"akuntan/helper"
)

// Handler untuk mendapatkan statistik dari pemasukan
func GetIncomeStatistics(w http.ResponseWriter, r *http.Request) {
	// Contoh data pemasukan
	incomes := []float64{1000000, 2000000, 1500000, 3000000} // Ganti dengan data dari database jika diperlukan

	totalIncome := helper.CalculateTotal(incomes)
	averageIncome := helper.CalculateAverage(incomes)

	response := map[string]interface{}{
		"total_income":   helper.FormatCurrency(totalIncome),
		"average_income": helper.FormatCurrency(averageIncome),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
