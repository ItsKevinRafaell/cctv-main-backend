package user

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

	// !!!!! INI BARIS DEBUGGING YANG PALING PENTING !!!!!
	// Baris ini akan mencetak isi struct 'user' ke terminal Go.
	log.Printf("DEBUG: Menerima data registrasi: %+v\n", user)

	if user.Role == "" {
		user.Role = "user"
	}

	err := h.service.Register(&user)
	if err != nil {
		// Kita cetak juga error aslinya untuk melihat apa yang terjadi
		log.Printf("ERROR: Gagal register service: %v\n", err)
		http.Error(w, "Email sudah terdaftar atau terjadi kesalahan lain", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User berhasil didaftarkan untuk perusahaan terkait."))
}
