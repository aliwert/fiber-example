package models

import "time"

type Car struct {
	ID           int       `json:"id"`
	Brand        string    `json:"brand"`
	Model        string    `json:"model"`
	Year         int       `json:"year"`
	Price        float64   `json:"price"`
	LicensePlate string    `json:"license_plate"`
	Color        string    `json:"color"`
	Status       string    `json:"status"` // Available, Rented, Maintenance
	Mileage      int       `json:"mileage"`
	CreatedAt    time.Time `json:"created_at"`
}

type CarRequest struct {
	Brand        string  `json:"brand" validate:"required"`
	Model        string  `json:"model" validate:"required"`
	Year         int     `json:"year" validate:"required,min=1900"`
	Price        float64 `json:"price" validate:"required,gt=0"`
	LicensePlate string  `json:"license_plate" validate:"required"`
	Color        string  `json:"color" validate:"required"`
	Mileage      int     `json:"mileage" validate:"required,gte=0"`
}
