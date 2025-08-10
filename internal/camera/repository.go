package camera

import (
	"cctv-main-backend/internal/domain"
	"database/sql"
)

type Repository interface {
	CreateCamera(camera *domain.Camera) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateCamera(camera *domain.Camera) (int64, error) {
	var cameraID int64
	query := `INSERT INTO cameras (name, location, company_id) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, camera.Name, camera.Location, camera.CompanyID).Scan(&cameraID)
	if err != nil {
		return 0, err
	}
	return cameraID, nil
}
