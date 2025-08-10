package company

import "cctv-main-backend/internal/domain"

type Service interface {
	Create(company *domain.Company) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(company *domain.Company) (int64, error) {
	return s.repo.CreateCompany(company)
}
