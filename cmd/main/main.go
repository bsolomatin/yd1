package main

import (
	"demo/first/Yd/internal/transport"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/calculate", transport.CalculateHandler)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}