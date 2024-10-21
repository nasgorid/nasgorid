// router/router.go
package router

import (
    "github.com/gorilla/mux"
    "akuntan/handler/auth"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    // Route untuk registrasi user
    r.HandleFunc("/register", handler.RegisterUser).Methods("POST")
    return r
}
