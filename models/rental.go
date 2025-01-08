package models

import "time"

type Rental struct {
	ID         int       `json:"id"`
	CarID      int       `json:"car_id"`
	CustomerID int       `json:"customer_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	TotalCost  float64   `json:"total_cost"`
	Status     string    `json:"status"` // Active, Completed, Cancelled
	CreatedAt  time.Time `json:"created_at"`
}

type RentalRequest struct {
	CarID      int       `json:"car_id" validate:"required"`
	CustomerID int       `json:"customer_id" validate:"required"`
	StartDate  time.Time `json:"start_date" validate:"required"`
	EndDate    time.Time `json:"end_date" validate:"required"`
}
