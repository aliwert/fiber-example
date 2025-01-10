package routes

import (
	"github.com/aliwert/fiber-example/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// Cars routes
	cars := api.Group("/cars")
	cars.Get("/", handlers.GetCars)
	cars.Get("/:id", handlers.GetCar)
	cars.Post("/create", handlers.CreateCar)
	cars.Put("/update/:id", handlers.UpdateCar)
	cars.Delete("/delete/:id", handlers.DeleteCar)

	// Customers routes
	customers := api.Group("/customers")
	customers.Get("/", handlers.GetCustomers)
	customers.Get("/:id", handlers.GetCustomer)
	customers.Post("/create", handlers.CreateCustomer)
	customers.Put("/update/:id", handlers.UpdateCustomer)
	customers.Delete("/delete/:id", handlers.DeleteCustomer)

	// Maintenance routes
	maintenance := api.Group("/maintenance")
	maintenance.Get("/", handlers.GetMaintenanceRecords)
	maintenance.Get("/:id", handlers.GetMaintenanceRecord)
	maintenance.Post("/create", handlers.CreateMaintenanceRecord)
	maintenance.Delete("/delete/:id", handlers.DeleteMaintenanceRecord)

	// Rentals routes
	rentals := api.Group("/rentals")
	rentals.Get("/", handlers.GetRentals)
	rentals.Get("/:id", handlers.GetRental)
	rentals.Post("/create", handlers.CreateRental)
	rentals.Put("/update/:id/status", handlers.UpdateRentalStatus)
}
