package handlers

import (
	"github.com/aliwert/fiber-example/models"

	"github.com/aliwert/fiber-example/database"

	"github.com/gofiber/fiber/v2"
)

func GetRentals(c *fiber.Ctx) error {
	rows, err := database.DB.Query(`
		SELECT id, car_id, customer_id, start_date, end_date, total_cost, status, created_at 
		FROM rentals
	`)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch rentals"})
	}
	defer rows.Close()

	var rentals []models.Rental
	for rows.Next() {
		var rental models.Rental
		if err := rows.Scan(
			&rental.ID, &rental.CarID, &rental.CustomerID, &rental.StartDate,
			&rental.EndDate, &rental.TotalCost, &rental.Status, &rental.CreatedAt,
		); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to scan rental"})
		}
		rentals = append(rentals, rental)
	}

	return c.JSON(rentals)
}

func GetRental(c *fiber.Ctx) error {
	id := c.Params("id")
	var rental models.Rental

	err := database.DB.QueryRow(`
		SELECT id, car_id, customer_id, start_date, end_date, total_cost, status, created_at 
		FROM rentals WHERE id = $1
	`, id).Scan(
		&rental.ID, &rental.CarID, &rental.CustomerID, &rental.StartDate,
		&rental.EndDate, &rental.TotalCost, &rental.Status, &rental.CreatedAt,
	)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Rental not found"})
	}

	return c.JSON(rental)
}

func CreateRental(c *fiber.Ctx) error {
	req := new(models.RentalRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	// Calculate total cost (you might want to implement your own pricing logic)
	totalCost := 100.0 // Placeholder

	var rental models.Rental
	err := database.DB.QueryRow(`
		INSERT INTO rentals (car_id, customer_id, start_date, end_date, total_cost, status)
		VALUES ($1, $2, $3, $4, $5, 'Active')
		RETURNING id, car_id, customer_id, start_date, end_date, total_cost, status, created_at
	`, req.CarID, req.CustomerID, req.StartDate, req.EndDate, totalCost).Scan(
		&rental.ID, &rental.CarID, &rental.CustomerID, &rental.StartDate,
		&rental.EndDate, &rental.TotalCost, &rental.Status, &rental.CreatedAt,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create rental"})
	}

	// Update car status
	_, err = database.DB.Exec("UPDATE cars SET status = 'Rented' WHERE id = $1", req.CarID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update car status"})
	}

	return c.Status(201).JSON(rental)
}

func UpdateRentalStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	status := c.Query("status")

	if status != "Active" && status != "Completed" && status != "Cancelled" {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid status"})
	}

	var rental models.Rental
	err := database.DB.QueryRow(`
		UPDATE rentals 
		SET status = $1
		WHERE id = $2
		RETURNING id, car_id, customer_id, start_date, end_date, total_cost, status, created_at
	`, status, id).Scan(
		&rental.ID, &rental.CarID, &rental.CustomerID, &rental.StartDate,
		&rental.EndDate, &rental.TotalCost, &rental.Status, &rental.CreatedAt,
	)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Rental not found"})
	}

	// If rental is completed or cancelled, update car status to Available
	if status == "Completed" || status == "Cancelled" {
		_, err = database.DB.Exec("UPDATE cars SET status = 'Available' WHERE id = $1", rental.CarID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to update car status"})
		}
	}

	return c.JSON(rental)
}
