package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Struktur Data Job / Transaksi Masuk
type Job struct {
	ID     int
	Amount float64
}

// Struktur Data Hasil Proses Worker
type Result struct {
	JobID    int
	Status   string
	WorkerID int
}

// Fungsi Worker yang berjalan di background (Goroutine)
func Worker(workerID int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("[Worker #%d] Mulai memproses Transaksi ID: %d (Nilai: Rp%.2f)\n", workerID, job.ID, job.Amount)

		// Simulasi waktu proses API pihak ketiga (misal: acak 500ms - 1500ms)
		processingTime := time.Duration(rand.Intn(1000)+500) * time.Millisecond
		time.Sleep(processingTime)

		// Kirim hasil pemrosesan ke channel results
		results <- Result{
			JobID:    job.ID,
			Status:   "SUCCESS",
			WorkerID: workerID,
		}
	}
}

func main() {
	fmt.Println("--- Advanced Go Backend: Worker Pool & Concurrency Control ---")

	const numJobs = 6
	const numWorkers = 3

	// Membuat buffered channels untuk antrean job dan hasil
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// 1. Spawn / Nyalakan sejumlah worker (Goroutines) secara bersamaan
	for w := 1; w <= numWorkers; w++ {
		go Worker(w, jobs, results)
	}

	// 2. Masukkan daftar transaksi (Jobs) ke dalam antrean
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{
			ID:     j,
			Amount: float64(j * 150000),
		}
	}
	close(jobs) // Tutup channel jobs setelah semua dimasukkan agar worker tahu antrean habis

	// 3. Kumpulkan semua hasil dari para worker
	for a := 1; a <= numJobs; a++ {
		res := <-results
		fmt.Printf("--> [Main Router] Transaksi #%d SELESAI diproses oleh Worker #%d (Status: %s)\n",
			res.JobID, res.WorkerID, res.Status)
	}

	fmt.Println("--- Semua antrean transaksi berhasil diselesaikan dengan aman! ---")
}
