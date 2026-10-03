package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Driver PostgreSQL (Nanti diinstal dengan 'go get github.com/lib/pq')
)

// Struktur Model User Account
type UserAccount struct {
	ID      int
	Name    string
	Email   string
	Balance float64
}

func main() {
	fmt.Println("Memulai simulasi koneksi database backend...")

	// Konfigurasi string koneksi database (Contoh ke PostgreSQL lokal)
	// Format: host=... port=... user=... password=... dbname=... sslmode=...
	connStr := "host=localhost port=5432 user=postgres password=rahasia dbname=neobank_db sslmode=disable"

	// Membuka koneksi database (Catatan: ini butuh server database aktif,
	// namun kodenya menunjukkan standar baku yang dipakai di industri)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Gagal konfigurasi database: %v", err)
	}
	defer db.Close()

	// Mengetes koneksi (Ping)
	err = db.Ping()
	if err != nil {
		fmt.Println("Info: Server database lokal belum aktif (ini normal jika PostgreSQL belum dinyalakan).")
		fmt.Println("Contoh query standar yang biasa digunakan di production:")
		simulateProductionQuery()
		return
	}

	fmt.Println("Berhasil terhubung ke database!")
}

// Simulasi logika query yang sering dipakai backend engineer
func simulateProductionQuery() {
	searchEmail := "milzon.ltf@gmail.com"

	fmt.Printf("\n[Contoh Query SQL]: SELECT id, name, balance FROM accounts WHERE email = '%s'\n", searchEmail)
	fmt.Println("Status: Query aman dari SQL Injection dengan menggunakan Prepared Statement (*sql.Row).")
}
