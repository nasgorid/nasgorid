package routes

import (
    "nasgorid_be/controller/auth"
    "nasgorid_be/controller/menu" // Tambahkan ini untuk menghubungkan controller menu
    "nasgorid_be/controller/pelanggan" // Import controller pelanggan
    "github.com/gorilla/mux"
)

// SetAuthRoutes defines routes for user authentication (register, login)
func SetAuthRoutes(router *mux.Router) {
    router.HandleFunc("/register", auth.Register).Methods("POST")
}

// SetMenuRoutes defines routes for menu operations (CRUD)
func SetMenuRoutes(router *mux.Router) {
    router.HandleFunc("/menu", menu.GetAllMenu).Methods("GET")          // Get all menus
    router.HandleFunc("/menu", menu.InsertMenu).Methods("POST")         // Insert a new menu
    router.HandleFunc("/menu/{id}", menu.UpdateMenu).Methods("PUT")     // Update menu by ID
    router.HandleFunc("/menu/{id}", menu.DeleteMenu).Methods("DELETE")  // Delete menu by ID
}

// SetPelangganRoutes defines routes for pelanggan operations
func SetPelangganRoutes(router *mux.Router) {
    router.HandleFunc("/pelanggan/{id}", pelanggan.GetPelangganByID).Methods("GET")     // Get pelanggan by ID
    router.HandleFunc("/pelanggan/{id}", pelanggan.UpdatePelanggan).Methods("PUT")      // Update pelanggan by ID
}

// InitializeRoutes combines all routes
func InitializeRoutes() *mux.Router {
    router := mux.NewRouter()

    // Authentication routes
    SetAuthRoutes(router)

    // Menu routes
    SetMenuRoutes(router)

    return router
}
