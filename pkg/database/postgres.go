package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// NewConnection membuat dan mengembalikan koneksi ke database PostgreSQL.
func NewConnection() *sql.DB {
	connStr := "host=127.0.0.1 port=5432 user=admin password=secret dbname=cctv_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Tidak bisa ping ke database: %v", err)
	}

	log.Println("✅ Berhasil terhubung ke database PostgreSQL!")
	return db
}

func Migrate(db *sql.DB) {
	createAnomalyTableQuery := `
	CREATE TABLE IF NOT EXISTS anomaly_reports (
		id SERIAL PRIMARY KEY,
		camera_id VARCHAR(50) NOT NULL,
		anomaly_type VARCHAR(50) NOT NULL,
		confidence FLOAT NOT NULL,
		reported_at TIMESTAMP WITH TIME ZONE NOT NULL
	);`
	if _, err := db.Exec(createAnomalyTableQuery); err != nil {
		log.Fatalf("Gagal membuat tabel anomaly_reports: %v", err)
	}
	log.Println("   > Tabel 'anomaly_reports' siap digunakan.")

	createUserTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	);`
	if _, err := db.Exec(createUserTableQuery); err != nil {
		log.Fatalf("Gagal membuat tabel users: %v", err)
	}
	log.Println("   > Tabel 'users' siap digunakan.")
}
