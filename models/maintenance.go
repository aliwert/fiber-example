package models

import "time"

type Maintenance struct {
	ID          int       `json:"id"`
	CarID       int       `json:"car_id"`
	ServiceType string    `json:"service_type"`
	Description string    `json:"description"`
	Cost        float64   `json:"cost"`
	ServiceDate time.Time `json:"service_date"`
	NextService time.Time `json:"next_service"`
	CreatedAt   time.Time `json:"created_at"`
}

type MaintenanceRequest struct {
	CarID       int       `json:"car_id" validate:"required"`
	ServiceType string    `json:"service_type" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Cost        float64   `json:"cost" validate:"required"`
	ServiceDate time.Time `json:"service_date" validate:"required"`
	NextService time.Time `json:"next_service" validate:"required"`
}
