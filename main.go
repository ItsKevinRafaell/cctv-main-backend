package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AnomalyReport struct {
	ID          int64     `json:"id"`
	CameraID    string    `json:"camera_id"`
	AnomalyType string    `json:"anomaly_type"`
	Confidence  float64   `json:"confidence"`
	ReportedAt  time.Time `json:"reported_at"`
}

var db *sql.DB

func reportAnomalyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var report AnomalyReport
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		http.Error(w, "Request body tidak valid", http.StatusBadRequest)
		return
	}

	log.Printf("✅ Laporan Diterima: Kamera %s, Tipe %s, Kepercayaan %.2f",
		report.CameraID, report.AnomalyType, report.Confidence)

	query := `
		INSERT INTO anomaly_reports (camera_id, anomaly_type, confidence, reported_at)
		VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(query, report.CameraID, report.AnomalyType, report.Confidence, time.Now())
	if err != nil {
		log.Printf("❌ Gagal menyimpan ke database: %v", err)
		http.Error(w, "Gagal memproses laporan", http.StatusInternalServerError)
		return
	}

	log.Println("   > Laporan berhasil disimpan ke database.")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Laporan berhasil diterima dan disimpan."))
}

func getAnomaliesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT id, camera_id, anomaly_type, confidence, reported_at FROM anomaly_reports ORDER BY reported_at DESC")
	if err != nil {
		log.Printf("❌ Gagal mengambil data dari database: %v", err)
		http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	reports := []AnomalyReport{}
	for rows.Next() {
		var report AnomalyReport
		if err := rows.Scan(&report.ID, &report.CameraID, &report.AnomalyType, &report.Confidence, &report.ReportedAt); err != nil {
			log.Printf("❌ Gagal memindai baris data: %v", err)
			continue
		}
		reports = append(reports, report)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Request body tidak valid", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Gagal memproses password", http.StatusInternalServerError)
		return
	}

	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2)`
	_, err = db.Exec(query, user.Email, string(hashedPassword))
	if err != nil {
		log.Printf("Gagal menyimpan user: %v", err)
		http.Error(w, "Email sudah terdaftar", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User berhasil didaftarkan."))
}

func main() {
	connStr := "user=admin password=secret dbname=cctv_db sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Tidak bisa ping ke database:", err)
	}
	log.Println("✅ Berhasil terhubung ke database PostgreSQL!")

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS anomaly_reports (
		id SERIAL PRIMARY KEY,
		camera_id VARCHAR(50) NOT NULL,
		anomaly_type VARCHAR(50) NOT NULL,
		confidence FLOAT NOT NULL,
		reported_at TIMESTAMP WITH TIME ZONE NOT NULL
	);`

	if _, err := db.Exec(createTableQuery); err != nil {
		log.Fatal("Gagal membuat tabel:", err)
	}
	log.Println("   > Tabel 'anomaly_reports' siap digunakan.")

	createUserTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	);`
	if _, err := db.Exec(createUserTableQuery); err != nil {
		log.Fatal("Gagal membuat tabel users:", err)
	}
	log.Println("   > Tabel 'users' siap digunakan.")

	http.HandleFunc("/api/report-anomaly", reportAnomalyHandler)
	http.HandleFunc("/api/anomalies", getAnomaliesHandler)
	http.HandleFunc("/api/register", registerHandler)

	port := "8080"
	fmt.Printf("Server backend Go berjalan di http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Gagal memulai server:", err)
	}
}
