// router/router.go
package router

import (
    "github.com/gorilla/mux"
    "akuntan/handler/auth"
    "akuntan/handler/produk"
    "akuntan/handler/transaksi_penjualan"
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
