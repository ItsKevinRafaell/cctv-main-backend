package anomaly

import (
	"cctv-main-backend/internal/domain"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	service Service
}

type ContextKey string

const UserClaimsKey = ContextKey("userClaims")

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateReport(w http.ResponseWriter, r *http.Request) {
	var report domain.AnomalyReport
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		http.Error(w, "Request body tidak valid", http.StatusBadRequest)
		return
	}

	err := h.service.SaveReport(&report)
	if err != nil {
		http.Error(w, "Gagal memproses laporan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Laporan berhasil diterima dan disimpan."))
}

func (h *Handler) GetAllReports(w http.ResponseWriter, r *http.Request) {
	// Retrieve claims from context using the same key
	claims, ok := r.Context().Value(UserClaimsKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "Gagal mengambil data pengguna dari token", http.StatusUnauthorized)
		fmt.Println("Failed to get claims from context.")
		return
	}

	fmt.Println("Token claims:", claims)

	companyID, ok := claims["company_id"].(float64)
	if !ok {
		http.Error(w, "Gagal mengambil company_id dari token", http.StatusUnauthorized)
		fmt.Println("Failed to extract company_id from claims.")
		return
	}

	reports, err := h.service.FetchAllReportsByCompany(int64(companyID))
	if err != nil {
		http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}
