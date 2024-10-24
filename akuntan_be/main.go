package main

import (
	"akuntan/config"
	"akuntan/handler/auth"
	"akuntan/handler/laporan"
	"akuntan/handler/pelanggan"
	"akuntan/handler/produk"
	"akuntan/handler/transaksi_pengeluaran"
	"akuntan/handler/transaksi_penjualan"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	config.InitMongoDB()
	fmt.Println("Hello World")

	router := mux.NewRouter()

	// Route untuk registrasi user
	router.HandleFunc("/register", auth.RegisterUser).Methods("POST")

	// Route untuk login user
	router.HandleFunc("/login", auth.LoginUser).Methods("POST")

	// Route untuk Produk
	router.HandleFunc("/products", produk.CreateProduct).Methods("POST")
	router.HandleFunc("/products", produk.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", produk.GetProductByID).Methods("GET")
	router.HandleFunc("/products/{id}", produk.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", produk.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/products-export-csv", produk.ExportProductsToCSV).Methods("GET")


	// Rute untuk transaksi penjualan
	router.HandleFunc("/transaksi", transaksi_penjualan.CreateSalesTransaction).Methods("POST")
	router.HandleFunc("/transaksi", transaksi_penjualan.GetSalesTransactions).Methods("GET")
	router.HandleFunc("/transaksi/{id}", transaksi_penjualan.GetSalesTransactionByID).Methods("GET")
	router.HandleFunc("/transaksi/{id}", transaksi_penjualan.UpdateSalesTransaction).Methods("PUT")
	router.HandleFunc("/transaksi/{id}", transaksi_penjualan.DeleteSalesTransaction).Methods("DELETE")
	router.HandleFunc("/transaksi-export-csv", transaksi_penjualan.ExportSalesTransactionsCSV).Methods("GET")

	// Rute untuk transaksi pengeluaran
	router.HandleFunc("/expense", transaksi_pengeluaran.CreateExpenseTransaction).Methods("POST")
	router.HandleFunc("/expense", transaksi_pengeluaran.GetExpenses).Methods("GET")
	router.HandleFunc("/expense/{id}", transaksi_pengeluaran.GetExpenseByID).Methods("GET")
	router.HandleFunc("/expense/{id}", transaksi_pengeluaran.UpdateExpense).Methods("PUT")
	router.HandleFunc("/expense/{id}", transaksi_pengeluaran.DeleteExpense).Methods("DELETE")

	// Customer Routes
	router.HandleFunc("/customers", pelanggan.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers", pelanggan.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", pelanggan.GetCustomerByID).Methods("GET")
	router.HandleFunc("/customers/{id}", pelanggan.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", pelanggan.DeleteCustomer).Methods("DELETE")

	// Rute untuk Laporan Keuangan
	router.HandleFunc("/reports", laporan.CreateFinancialReport).Methods("POST")
	router.HandleFunc("/reports", laporan.GetFinancialReports).Methods("GET")
	router.HandleFunc("/reports/{id}", laporan.GetFinancialReportByID).Methods("GET")
	router.HandleFunc("/reports/{id}", laporan.DeleteFinancialReport).Methods("DELETE")

	allowedOrigins := []string{"http://127.0.0.1:5500"}

	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		Debug:            true,
	})

	handler := c.Handler(router)

	fmt.Println("Server is running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", handler))

}
