package handlers

import (
	"log"

	"github.com/aliwert/fiber-example/database"
	"github.com/aliwert/fiber-example/models"
	"github.com/gofiber/fiber/v2"
)

func GetCars(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, brand, model, year, price, created_at FROM cars")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch cars"})
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ID, &car.Brand, &car.Model, &car.Year, &car.Price, &car.CreatedAt); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to scan car"})
		}
		cars = append(cars, car)
	}

	return c.JSON(cars)
}

func GetCar(c *fiber.Ctx) error {
	id := c.Params("id")
	var car models.Car

	err := database.DB.QueryRow("SELECT id, brand, model, year, price, created_at FROM cars WHERE id = $1", id).
		Scan(&car.ID, &car.Brand, &car.Model, &car.Year, &car.Price, &car.CreatedAt)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Car not found"})
	}

	return c.JSON(car)
}

func CreateCar(c *fiber.Ctx) error {
	car := new(models.Car)
	if err := c.BodyParser(car); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	// Validate required fields
	if car.Brand == "" || car.Model == "" || car.Year == 0 ||
		car.Price == 0 || car.LicensePlate == "" || car.Color == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Missing required fields",
			"required_fields": []string{
				"brand", "model", "year", "price", "license_plate", "color", "mileage",
			},
		})
	}

	query := `
        INSERT INTO cars (brand, model, year, price, license_plate, color, mileage, status)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id, brand, model, year, price, license_plate, color, mileage, status, created_at
    `

	err := database.DB.QueryRow(
		query,
		car.Brand,
		car.Model,
		car.Year,
		car.Price,
		car.LicensePlate,
		car.Color,
		car.Mileage,
		car.Status,
	).Scan(
		&car.ID,
		&car.Brand,
		&car.Model,
		&car.Year,
		&car.Price,
		&car.LicensePlate,
		&car.Color,
		&car.Mileage,
		&car.Status,
		&car.CreatedAt,
	)

	if err != nil {
		// Log the actual error
		log.Printf("Database error: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(car)
}

func UpdateCar(c *fiber.Ctx) error {
	id := c.Params("id")
	car := new(models.Car)

	if err := c.BodyParser(car); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	query := `
		UPDATE cars 
		SET brand = $1, model = $2, year = $3, price = $4
		WHERE id = $5
		RETURNING id, brand, model, year, price, created_at
	`

	err := database.DB.QueryRow(query, car.Brand, car.Model, car.Year, car.Price, id).
		Scan(&car.ID, &car.Brand, &car.Model, &car.Year, &car.Price, &car.CreatedAt)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Car not found"})
	}

	return c.JSON(car)
}

func DeleteCar(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := database.DB.Exec("DELETE FROM cars WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete car"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Car not found"})
	}
	return c.SendStatus(204)
}
