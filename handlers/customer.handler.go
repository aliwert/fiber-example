package handlers

import (
	"github.com/aliwert/fiber-example/database"
	"github.com/aliwert/fiber-example/models"

	"github.com/gofiber/fiber/v2"
)

func GetCustomers(c *fiber.Ctx) error {
	rows, err := database.DB.Query(`
		SELECT id, first_name, last_name, email, phone, address, created_at 
		FROM customers
	`)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch customers"})
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		if err := rows.Scan(
			&customer.ID, &customer.FirstName, &customer.LastName,
			&customer.Email, &customer.Phone, &customer.Address, &customer.CreatedAt,
		); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to scan customer"})
		}
		customers = append(customers, customer)
	}

	return c.JSON(customers)
}

func GetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer models.Customer

	err := database.DB.QueryRow(`
		SELECT id, first_name, last_name, email, phone, address, created_at 
		FROM customers WHERE id = $1
	`, id).Scan(
		&customer.ID, &customer.FirstName, &customer.LastName,
		&customer.Email, &customer.Phone, &customer.Address, &customer.CreatedAt,
	)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Customer not found"})
	}

	return c.JSON(customer)
}

func CreateCustomer(c *fiber.Ctx) error {
	req := new(models.CustomerRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	var customer models.Customer
	err := database.DB.QueryRow(`
		INSERT INTO customers (first_name, last_name, email, phone, address)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, first_name, last_name, email, phone, address, created_at
	`, req.FirstName, req.LastName, req.Email, req.Phone, req.Address).Scan(
		&customer.ID, &customer.FirstName, &customer.LastName,
		&customer.Email, &customer.Phone, &customer.Address, &customer.CreatedAt,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create customer"})
	}

	return c.Status(201).JSON(customer)
}

func UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	req := new(models.CustomerRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	var customer models.Customer
	err := database.DB.QueryRow(`
		UPDATE customers 
		SET first_name = $1, last_name = $2, email = $3, phone = $4, address = $5
		WHERE id = $6
		RETURNING id, first_name, last_name, email, phone, address, created_at
	`, req.FirstName, req.LastName, req.Email, req.Phone, req.Address, id).Scan(
		&customer.ID, &customer.FirstName, &customer.LastName,
		&customer.Email, &customer.Phone, &customer.Address, &customer.CreatedAt,
	)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Customer not found"})
	}

	return c.JSON(customer)
}

func DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := database.DB.Exec("DELETE FROM customers WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete customer"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Customer not found"})
	}

	return c.SendStatus(204)
}
