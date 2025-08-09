package anomaly

import (
	"cctv-main-backend/internal/domain"
	"database/sql"
	"time"
)

type Repository interface {
	CreateReport(report *domain.AnomalyReport) error
	GetAllReports() ([]domain.AnomalyReport, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateReport(report *domain.AnomalyReport) error {
	query := `INSERT INTO anomaly_reports (camera_id, anomaly_type, confidence, reported_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, report.CameraID, report.AnomalyType, report.Confidence, time.Now())
	return err
}

func (r *repository) GetAllReports() ([]domain.AnomalyReport, error) {
	rows, err := r.db.Query("SELECT id, camera_id, anomaly_type, confidence, reported_at FROM anomaly_reports ORDER BY reported_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []domain.AnomalyReport
	for rows.Next() {
		var report domain.AnomalyReport
		if err := rows.Scan(&report.ID, &report.CameraID, &report.AnomalyType, &report.Confidence, &report.ReportedAt); err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}
	return reports, nil
}
