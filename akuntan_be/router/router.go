// router/router.go
package router

import (
    "github.com/gorilla/mux"
    "akuntan/handler/auth"
    "akuntan/handler/produk"
    "akuntan/handler/transaksi_penjualan"
    "akuntan/handler/transaksi_pengeluaran"
    "akuntan/handler/pelanggan"
    "net/http"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    // Route untuk registrasi user
    r.HandleFunc("/register", auth.RegisterUser).Methods("POST")

    // Route untuk login user
    r.HandleFunc("/login", auth.LoginUser).Methods("POST")


    	// Route untuk Produk
	r.HandleFunc("/products", produk.CreateProduct).Methods("POST")
	r.HandleFunc("/products", produk.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", produk.GetProductByID).Methods("GET")
	r.HandleFunc("/products/{id}", produk.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", produk.DeleteProduct).Methods("DELETE")

    // Rute untuk transaksi penjualan
	r.HandleFunc("/transaksi", transaksi_penjualan.CreateSalesTransaction).Methods("POST")
	r.HandleFunc("/transaksi", transaksi_penjualan.GetSalesTransactions).Methods("GET")
	r.HandleFunc("/transaksi/{id}", transaksi_penjualan.GetSalesTransactionByID).Methods("GET")
	r.HandleFunc("/transaksi/{id}", transaksi_penjualan.UpdateSalesTransaction).Methods("PUT")
	r.HandleFunc("/transaksi/{id}", transaksi_penjualan.DeleteSalesTransaction).Methods("DELETE")
    r.HandleFunc("/transaksi/export-csv", transaksi_penjualan.ExportSalesTransactionsCSV).Methods("GET")


     // Rute untuk transaksi pengeluaran
    r.HandleFunc("/expense", transaksi_pengeluaran.CreateExpenseTransaction).Methods("POST")
    r.HandleFunc("/expense", transaksi_pengeluaran.GetExpenses).Methods("GET")
    r.HandleFunc("/expense/{id}", transaksi_pengeluaran.GetExpenseByID).Methods("GET")
    r.HandleFunc("/expense/{id}", transaksi_pengeluaran.UpdateExpense).Methods("PUT")
    r.HandleFunc("/expense/{id}", transaksi_pengeluaran.DeleteExpense).Methods("DELETE")

    	// Customer Routes
	r.HandleFunc("/customers", pelanggan.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers", pelanggan.GetCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", pelanggan.GetCustomerByID).Methods("GET")
	r.HandleFunc("/customers/{id}", pelanggan.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", pelanggan.DeleteCustomer).Methods("DELETE")



    // // Jika ada route yang memerlukan autentikasi
    // protected := r.PathPrefix("/protected").Subrouter()
    // protected.Use(auth.AuthMiddleware)
    // protected.HandleFunc("/", ProtectedHandler).Methods("GET")
 return r
}

// Contoh handler untuk route yang dilindungi
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
    email := r.Context().Value("email").(string)
    w.Write([]byte("Welcome to the protected area, " + email))
}
