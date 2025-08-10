package domain

import "time"

type AnomalyReport struct {
	ID          int64     `json:"id"`
	CameraID    int64     `json:"camera_id"`
	AnomalyType string    `json:"anomaly_type"`
	Confidence  float64   `json:"confidence"`
	ReportedAt  time.Time `json:"reported_at"`
}
