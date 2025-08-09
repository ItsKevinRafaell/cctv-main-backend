package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AnomalyReport struct {
	CameraID    string  `json:"camera_id"`
	AnomalyType string  `json:"anomaly_type"`
	Confidence  float64 `json:"confidence"`
}

func reportAnomalyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var report AnomalyReport

	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, "Request body tidak valid", http.StatusBadRequest)
		return
	}

	log.Printf("✅ Laporan Anomali Diterima!")
	log.Printf("   > Kamera: %s", report.CameraID)
	log.Printf("   > Tipe: %s", report.AnomalyType)
	log.Printf("   > Kepercayaan: %.2f", report.Confidence)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Laporan berhasil diterima."))
}

func main() {
	http.HandleFunc("/api/report-anomaly", reportAnomalyHandler)

	port := "8080"
	fmt.Printf("Server backend Go berjalan di http://localhost:%s\n", port)
	fmt.Println("Menunggu laporan di endpoint /api/report-anomaly...")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Gagal memulai server:", err)
	}
}