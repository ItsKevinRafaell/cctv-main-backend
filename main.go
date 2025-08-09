package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type AnomalyReport struct {
	CameraID    string  `json:"camera_id"`
	AnomalyType string  `json:"anomaly_type"`
	Confidence  float64 `json:"confidence"`
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

	http.HandleFunc("/api/report-anomaly", reportAnomalyHandler)

	port := "8080"
	fmt.Printf("Server backend Go berjalan di http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Gagal memulai server:", err)
	}
}
