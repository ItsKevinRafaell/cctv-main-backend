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

	if user.Role == "" {
		user.Role = "user"
	}

	err := h.service.Register(&user)
	if err != nil {
		http.Error(w, "Email sudah terdaftar atau terjadi kesalahan lain", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User berhasil didaftarkan untuk perusahaan terkait."))
}

// curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55X2lkIjoxLCJlbWFpbCI6ImJ1ZGlAamF5YWFiYWRpLmNvbSIsImV4cCI6MTc1NTA2ODQ4MSwicm9sZSI6ImNvbXBhbnlfYWRtaW4iLCJ1c2VyX2lkIjoyfQ.gXS1PKtuUmQ0nFC9NQ_OqJKC0rdV81Ar5EB3w9W-E3E" http://localhost:8080/api/anomalies
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55X2lkIjoxLCJlbWFpbCI6ImJ1ZGlAamF5YWFiYWRpLmNvbSIsImV4cCI6MTc1NTA2ODI1NCwicm9sZSI6ImNvbXBhbnlfYWRtaW4iLCJ1c2VyX2lkIjoyfQ.9hLEUQUpvSlboN_chI49Ke0Pwtfrh5QYSO-pJhEoRW8

// curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55X2lkIjoxLCJlbWFpbCI6ImJ1ZGlAamF5YWFiYWRpLmNvbSIsImV4cCI6MTc1NTA2ODY4MCwicm9sZSI6ImNvbXBhbnlfYWRtaW4iLCJ1c2VyX2lkIjoyfQ.RdHspdIuLg-XKsxHFRDZ18q1UJrr-VBMDNMtRb0XbpQ" http://localhost:8080/api/anomalies

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55X2lkIjozLCJlbWFpbCI6ImRpbmFAbWFrbXVyLmNvbSIsImV4cCI6MTc1NTA2ODI3OSwicm9sZSI6InVzZXIiLCJ1c2VyX2lkIjo0fQ.B63X6r8KX9dnDiNDrLpUr7TbSGf4mvyNR6PyBj0WkN0
