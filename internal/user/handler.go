package user

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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var input domain.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Request body tidak valid", http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Request body tidak valid", http.StatusBadRequest)
		return
	}

	err := h.service.Register(&user)
	if err != nil {
		http.Error(w, "Email sudah terdaftar", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User berhasil didaftarkan."))
}
