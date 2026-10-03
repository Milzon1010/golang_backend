package main

import (
	"fmt"
	"time"
)

// Simulasi fungsi pengecekan transaksi ke partner perbankan secara asynchronous
func CheckTransactionGateway(gatewayName string, txID string, ch chan string) {
	fmt.Printf("[Worker] Memulai pengecekan %s untuk transaksi %s...\n", gatewayName, txID)

	// Simulasi waktu proses jaringan (misal: 1 detik)
	time.Sleep(1 * time.Second)

	result := fmt.Sprintf("Transaksi %s via %s: BERHASIL (Verified)", txID, gatewayName)

	// Mengirim hasil ke channel
	ch <- result
}

func main() {
	fmt.Println("--- Memulai Simulasi Backend Concurrency (Goroutines) ---")

	// Membuat channel untuk komunikasi antar goroutines
	ch := make(chan string)

	// Daftar gateway pembayaran yang dicek secara bersamaan (paralel)
	start := time.Now()

	go CheckTransactionGateway("Payment-Gateway-A", "TX-9901", ch)
	go CheckTransactionGateway("Payment-Gateway-B", "TX-9902", ch)
	go CheckTransactionGateway("AI-Fraud-Detector", "TX-9903", ch)

	// Menangkap hasil dari ketiga goroutines melalui channel
	for i := 0; i < 3; i++ {
		response := <-ch
		fmt.Println("[Hasil Sistem]:", response)
	}

	duration := time.Since(start)
	fmt.Printf("\nSemua proses paralel selesai dalam waktu: %v\n", duration)
	fmt.Println("Catatan: Proses berjalan serentak, jauh lebih cepat dibanding berurutan (sequential)!")
}
