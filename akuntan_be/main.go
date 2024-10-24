package main

import (
	"akuntan/config"
	"akuntan/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.InitMongoDB()
	fmt.Println("Hello World")

	r := router.SetupRouter()
	handler := router.SetupCORS(r)

	fmt.Println("Server is running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", handler))
}
