package anomaly

import (
	"cctv-main-backend/internal/domain"
	"database/sql"
	"time"
)

type Repository interface {
	CreateReport(report *domain.AnomalyReport) error
	GetAllReportsByCompany(companyID int64) ([]domain.AnomalyReport, error)
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
func (r *repository) GetAllReportsByCompany(companyID int64) ([]domain.AnomalyReport, error) {
	query := `
		SELECT r.id, r.camera_id, r.anomaly_type, r.confidence, r.reported_at
		FROM anomaly_reports r
		JOIN cameras c ON r.camera_id = c.id
		WHERE c.company_id = $1
		ORDER BY r.reported_at DESC`

	rows, err := r.db.Query(query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []domain.AnomalyReport
	for rows.Next() {
		var report domain.AnomalyReport
		// Sesuaikan Scan karena kita tidak mengambil semua kolom dari JOIN
		if err := rows.Scan(&report.ID, &report.CameraID, &report.AnomalyType, &report.Confidence, &report.ReportedAt); err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}
	return reports, nil
}
