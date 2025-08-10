package anomaly

import "cctv-main-backend/internal/domain"

type Service interface {
	SaveReport(report *domain.AnomalyReport) error
	FetchAllReportsByCompany(companyID int64) ([]domain.AnomalyReport, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) SaveReport(report *domain.AnomalyReport) error {
	// Di masa depan, logika bisnis kompleks bisa ditambahkan di sini
	return s.repo.CreateReport(report)
}

func (s *service) FetchAllReportsByCompany(companyID int64) ([]domain.AnomalyReport, error) {
	return s.repo.GetAllReportsByCompany(companyID)
}
