package main

import (
	"cctv-main-backend/internal/anomaly"
	"cctv-main-backend/internal/user"
	"cctv-main-backend/pkg/database"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database.NewConnection()
	database.Migrate(db)
	defer db.Close()

	// Inisialisasi Departemen Anomali (Dependency Injection)
	anomalyRepo := anomaly.NewRepository(db)
	anomalyService := anomaly.NewService(anomalyRepo)
	anomalyHandler := anomaly.NewHandler(anomalyService)

	// Inisialisasi Departemen User (Dependency Injection)
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	http.HandleFunc("/api/register", userHandler.Register)
	http.HandleFunc("/api/login", userHandler.Login)
	http.HandleFunc("/api/report-anomaly", anomalyHandler.CreateReport)
	http.HandleFunc("/api/anomalies", anomalyHandler.GetAllReports)

	port := "8080"
	fmt.Printf("Server berjalan di http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Gagal memulai server:", err)
	}
}
