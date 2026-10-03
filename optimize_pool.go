package main

import (
	"bytes"
	"fmt"
	"sync"
)

// Menyiapkan sync.Pool untuk mendaur ulang buffer memori
var bufPool = sync.Pool{
	New: func() interface{} {
		// Alokasi memori awal untuk buffer string/json response
		return new(bytes.Buffer)
	},
}

func processApiResponse(dataID int) string {
	// 1. Ambil buffer yang sudah tersedia dari pool (menghemat alokasi memori baru)
	buf := bufPool.Get().(*bytes.Buffer)

	// Pastikan buffer dibersihkan setelah selesai digunakan
	buf.Reset()
	defer bufPool.Put(buf)

	// 2. Menulis data respons API ke dalam buffer
	buf.WriteString("{\"status\":\"success\",\"transaction_id\":")
	buf.WriteString(fmt.Sprintf("%d", dataID))
	buf.WriteString(",\"message\":\"Processed via memory pool optimization\"}")

	return buf.String()
}

func main() {
	fmt.Println("--- Advanced Go Backend: Performance Optimization (sync.Pool) ---")

	// Simulasi pemrosesan ribuan request dengan memori yang di-recycle
	for i := 1; i <= 5; i++ {
		response := processApiResponse(9000 + i)
		fmt.Printf("[Response %d]: %s\n", i, response)
	}

	fmt.Println("\n[Info Optimasi]: Dengan sync.Pool, beban Garbage Collector berkurang drastis pada sistem high-throughput!")
}
