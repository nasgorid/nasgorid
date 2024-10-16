package routes

import (
    "nasgorid_be/controller/auth"

    "github.com/gorilla/mux"
)

// SetAuthRoutes defines routes for user authentication (register, login)
func SetAuthRoutes(router *mux.Router) {
    router.HandleFunc("/register", auth.Register).Methods("POST")
}