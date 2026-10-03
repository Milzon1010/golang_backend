package main

import (
	"errors"
	"fmt"
)

type TransactionRequest struct {
	SourceAccount      string  `json:"source_account"`
	DestinationAccount string  `json:"destination_account"`
	Amount             float64 `json:"amount"`
}

// Fungsi simulasi transfer dengan validasi ketat
func ProcessTransfer(req TransactionRequest, currentBalance float64) error {
	if req.Amount <= 0 {
		return errors.New("validasi gagal: jumlah transfer harus lebih besar dari 0")
	}
	if req.Amount > currentBalance {
		return errors.New("transaksi ditolak: saldo tidak mencukupi")
	}
	return nil
}

func main() {
	// Simulasi pengujian transaksi
	tx := TransactionRequest{
		SourceAccount:      "ACC-001",
		DestinationAccount: "ACC-002",
		Amount:             6000000.0, // Mencoba transfer lebih besar dari saldo
	}

	balance := 5000000.0

	fmt.Printf("Memproses transfer sebesar Rp%.2f...\n", tx.Amount)
	err := ProcessTransfer(tx, balance)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sukses: Transaksi berhasil diproses.")
	}
}
