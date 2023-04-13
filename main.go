package main

import (
	"fmt"
	"net/http"
	"go-person-api/controllers"
)

func main() {
	http.HandleFunc("/", controllers.HomeHandler)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}