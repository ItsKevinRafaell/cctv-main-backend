package company

import (
	"cctv-main-backend/internal/domain"
	"encoding/json"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company domain.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, "Request body tidak valid", http.StatusBadRequest)
		return
	}

	companyID, err := h.service.Create(&company)
	if err != nil {
		http.Error(w, "Gagal membuat perusahaan", http.StatusInternalServerError)
		return
	}

	response := map[string]int64{"company_id": companyID}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
