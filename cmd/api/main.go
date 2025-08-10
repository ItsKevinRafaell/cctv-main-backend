package main

import (
	"cctv-main-backend/internal/anomaly"
	"cctv-main-backend/internal/company"
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

	anomalyRepo := anomaly.NewRepository(db)
	anomalyService := anomaly.NewService(anomalyRepo)
	anomalyHandler := anomaly.NewHandler(anomalyService)

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	companyRepo := company.NewRepository(db)
	companyService := company.NewService(companyRepo)
	companyHandler := company.NewHandler(companyService)

	http.HandleFunc("/api/register", userHandler.Register)
	http.HandleFunc("/api/login", userHandler.Login)

	http.HandleFunc("/api/report-anomaly", anomalyHandler.CreateReport)
	http.HandleFunc("/api/anomalies", authMiddleware(anomalyHandler.GetAllReports))

	http.HandleFunc("/api/companies", companyHandler.CreateCompany)

	port := "8080"
	fmt.Printf("Server berjalan di http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Gagal memulai server:", err)
	}
}
