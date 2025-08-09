package anomaly

import (
	"cctv-main-backend/internal/domain"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateReport(w http.ResponseWriter, r *http.Request) {
	var report domain.AnomalyReport
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		http.Error(w, "Request body tidak valid", http.StatusBadRequest)
		return
	}

	log.Printf("✅ Laporan Diterima: Kamera %s", report.CameraID)
	err := h.service.SaveReport(&report)
	if err != nil {
		log.Printf("❌ Gagal memproses laporan: %v", err)
		http.Error(w, "Gagal memproses laporan", http.StatusInternalServerError)
		return
	}

	log.Println("   > Laporan berhasil disimpan.")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Laporan berhasil diterima dan disimpan."))
}

func (h *Handler) GetAllReports(w http.ResponseWriter, r *http.Request) {
	reports, err := h.service.FetchAllReports()
	if err != nil {
		http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}
