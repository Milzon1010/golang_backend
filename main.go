package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struktur data untuk Account
type Account struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

// Database simulasi di dalam memori
var accounts = []Account{
	{ID: "1", Name: "Milzon", Balance: 5000000.0},
	{ID: "2", Name: "Budi", Balance: 1500000.0},
}

func getAccountsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func main() {
	// Menyiapkan endpoint route
	http.HandleFunc("/api/accounts", getAccountsHandler)

	fmt.Println("Server berjalan di port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}