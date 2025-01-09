package handlers

import (
	"github.com/aliwert/fiber-example/database"
	"github.com/aliwert/fiber-example/models"

	"github.com/gofiber/fiber/v2"
)

func GetMaintenanceRecords(c *fiber.Ctx) error {
	rows, err := database.DB.Query(`
		SELECT id, car_id, service_type, description, cost, service_date, next_service, created_at 
		FROM maintenance
	`)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch maintenance records"})
	}
	defer rows.Close()

	var records []models.Maintenance
	for rows.Next() {
		var record models.Maintenance
		if err := rows.Scan(
			&record.ID, &record.CarID, &record.ServiceType, &record.Description,
			&record.Cost, &record.ServiceDate, &record.NextService, &record.CreatedAt,
		); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to scan maintenance record"})
		}
		records = append(records, record)
	}

	return c.JSON(records)
}

func GetMaintenanceRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	var record models.Maintenance

	err := database.DB.QueryRow(`
		SELECT id, car_id, service_type, description, cost, service_date, next_service, created_at 
		FROM maintenance WHERE id = $1
	`, id).Scan(
		&record.ID, &record.CarID, &record.ServiceType, &record.Description,
		&record.Cost, &record.ServiceDate, &record.NextService, &record.CreatedAt,
	)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Maintenance record not found"})
	}

	return c.JSON(record)
}

func CreateMaintenanceRecord(c *fiber.Ctx) error {
	req := new(models.MaintenanceRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	var record models.Maintenance
	err := database.DB.QueryRow(`
		INSERT INTO maintenance (car_id, service_type, description, cost, service_date, next_service)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, car_id, service_type, description, cost, service_date, next_service, created_at
	`, req.CarID, req.ServiceType, req.Description, req.Cost, req.ServiceDate, req.NextService).Scan(
		&record.ID, &record.CarID, &record.ServiceType, &record.Description,
		&record.Cost, &record.ServiceDate, &record.NextService, &record.CreatedAt,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create maintenance record"})
	}

	return c.Status(201).JSON(record)
}

func DeleteMaintenanceRecord(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := database.DB.Exec("DELETE FROM maintenance WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete maintenance record"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Maintenance record not found"})
	}

	return c.SendStatus(204)
}
