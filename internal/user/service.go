package user

import (
	"cctv-main-backend/internal/domain"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(user *domain.User) error
	Login(input *domain.User) (string, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Login(input *domain.User) (string, error) {
	user, err := s.repo.GetUserByEmail(input.Email)
	if err != nil {
		return "", errors.New("email atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		return "", errors.New("email atau password salah")
	}

	// 3. Jika berhasil, kita akan buat token di sini (untuk sekarang, kita kembalikan string sukses)
	// TODO: Implementasi pembuatan JWT Token

	return "login_berhasil_token_akan_dibuat_di_sini", nil
}

func (s *service) Register(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)
	return s.repo.CreateUser(user)
}
