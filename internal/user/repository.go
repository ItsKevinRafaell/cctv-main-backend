package user

import (
	"cctv-main-backend/internal/domain"
	"database/sql"
)

type Repository interface {
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, email, password_hash FROM users WHERE email=$1`

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) CreateUser(user *domain.User) error {
	query := `INSERT INTO users (email, password_hash, company_id, role) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, user.Email, user.PasswordHash, user.CompanyID, user.Role)
	return err
}
