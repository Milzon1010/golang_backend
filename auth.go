package main

import (
	"fmt"
	"time"
)

// Simulasi struktur klaim data pengguna di dalam token
type UserClaims struct {
	Username  string
	Role      string
	ExpiresAt int64
}

// Fungsi simulasi pembuatan Token (Token Generation)
func GenerateMockToken(username string, role string) string {
	// Menentukan waktu kedaluwarsa token (misal: 2 jam dari sekarang)
	expirationTime := time.Now().Add(2 * time.Hour).Unix()

	// Simulasi string token yang di-encode
	mockToken := fmt.Sprintf("header.payload_user_%s_role_%s.signature_hash_secret", username, role)

	fmt.Printf("[Auth Service] Token berhasil dibuat untuk %s (Role: %s)\n", username, role)
	fmt.Printf("[Auth Service] Expires at: %s\n", time.Unix(expirationTime, 0).Format(time.RFC3339))

	return mockToken
}

// Fungsi simulasi verifikasi Token (Token Validation Middleware)
func VerifyToken(token string) bool {
	if token != "" {
		// Di implementasi nyata, di sini akan dicek validitas signature dan masa aktifnya
		return true
	}
	return false
}

func main() {
	fmt.Println("--- Memulai Simulasi JWT Auth Backend ---")

	// 1. Pengguna melakukan login berhasil, lalu sistem membuatkan token
	token := GenerateMockToken("milzon_tech", "Admin_Backend")
	fmt.Println("Generated Token:", token)

	// 2. Pengguna mengakses API terproteksi dengan membawa token tersebut
	fmt.Println("\n[Middleware] Memverifikasi token pada request API...")
	isValid := VerifyToken(token)

	if isValid {
		fmt.Println("Akses Diterima: Selamat datang, Milzon! Silakan akses data API.")
	} else {
		fmt.Println("Akses Ditolak: Token tidak valid atau kedaluwarsa.")
	}
}
