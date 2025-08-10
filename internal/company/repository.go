package company

import (
	"cctv-main-backend/internal/domain"
	"database/sql"
)

type Repository interface {
	CreateCompany(company *domain.Company) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateCompany(company *domain.Company) (int64, error) {
	var companyID int64
	query := `INSERT INTO companies (name) VALUES ($1) RETURNING id`
	err := r.db.QueryRow(query, company.Name).Scan(&companyID)
	if err != nil {
		return 0, err
	}
	return companyID, nil
}
