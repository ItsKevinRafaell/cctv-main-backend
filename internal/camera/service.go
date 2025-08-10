package camera

import "cctv-main-backend/internal/domain"

type Service interface {
	RegisterCamera(camera *domain.Camera) (int64, error)
	GetCamerasForCompany(companyID int64) ([]domain.Camera, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) RegisterCamera(camera *domain.Camera) (int64, error) {
	return s.repo.CreateCamera(camera)
}

func (s *service) GetCamerasForCompany(companyID int64) ([]domain.Camera, error) {
	return s.repo.GetCamerasByCompanyID(companyID)
}
