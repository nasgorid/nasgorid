package router

import (
	"akuntan/handler/auth"
	"akuntan/handler/laporan"
	"akuntan/handler/pelanggan"
	"akuntan/handler/produk"
	"akuntan/handler/transaksi_pengeluaran"
	"akuntan/handler/transaksi_penjualan"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

// SetupRouter initializes the routes for the application.
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Route for user registration
	router.HandleFunc("/register", auth.RegisterUser).Methods("POST")

	// Route for user login
	router.HandleFunc("/login", auth.LoginUser).Methods("POST")

	// Product Routes
	router.HandleFunc("/products", produk.CreateProduct).Methods("POST")
	router.HandleFunc("/products", produk.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", produk.GetProductByID).Methods("GET")
	router.HandleFunc("/products/{id}", produk.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", produk.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/products-export-csv", produk.ExportProductsToCSV).Methods("GET")

	// Sales Transaction Routes
	router.HandleFunc("/transaksi", transaksi_penjualan.CreateSalesTransaction).Methods("POST")
	router.HandleFunc("/transaksi", transaksi_penjualan.GetSalesTransactions).Methods("GET")
	router.HandleFunc("/transaksi/{id}", transaksi_penjualan.GetSalesTransactionByID).Methods("GET")
	router.HandleFunc("/transaksi/{id}", transaksi_penjualan.UpdateSalesTransaction).Methods("PUT")
	router.HandleFunc("/transaksi/{id}", transaksi_penjualan.DeleteSalesTransaction).Methods("DELETE")
	router.HandleFunc("/transaksi-export-csv", transaksi_penjualan.ExportSalesTransactionsCSV).Methods("GET")

	// Expense Transaction Routes
	router.HandleFunc("/expense", transaksi_pengeluaran.CreateExpenseTransaction).Methods("POST")
	router.HandleFunc("/expense", transaksi_pengeluaran.GetExpenses).Methods("GET")
	router.HandleFunc("/expense/{id}", transaksi_pengeluaran.GetExpenseByID).Methods("GET")
	router.HandleFunc("/expense/{id}", transaksi_pengeluaran.UpdateExpense).Methods("PUT")
	router.HandleFunc("/expense/{id}", transaksi_pengeluaran.DeleteExpense).Methods("DELETE")
	router.HandleFunc("/expense-export-csv", transaksi_pengeluaran.ExportExpensesToCSV).Methods("GET")

	// Customer Routes
	router.HandleFunc("/customers", pelanggan.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers", pelanggan.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", pelanggan.GetCustomerByID).Methods("GET")
	router.HandleFunc("/customers/{id}", pelanggan.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", pelanggan.DeleteCustomer).Methods("DELETE")

	// Financial Report Routes
	router.HandleFunc("/reports", laporan.CreateFinancialReport).Methods("POST")
	router.HandleFunc("/reports", laporan.GetFinancialReports).Methods("GET")
	router.HandleFunc("/reports/{id}", laporan.GetFinancialReportByID).Methods("GET")
	router.HandleFunc("/reports/{id}", laporan.DeleteFinancialReport).Methods("DELETE")

	return router
}

// SetupCORS configures CORS for the router.
func SetupCORS(router *mux.Router) http.Handler {
	allowedOrigins := []string{"http://127.0.0.1:5500"}

	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		Debug:            true,
	})

	return c.Handler(router)
}
