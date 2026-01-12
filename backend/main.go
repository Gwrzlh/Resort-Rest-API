package main

import (
	"github.com/Gwrzlh/Resort-Rest-API/backend/config"
	"fmt"
	"net/http"
)

func main() {

	db := config.ConnectDB()
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "RES API is running with database connection!")
	})

	fmt.Println("Server running di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}