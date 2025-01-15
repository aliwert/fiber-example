package routes

import (
	"github.com/aliwert/fiber-example/handlers"
	"github.com/aliwert/fiber-example/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// Public routes (no auth required)
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	//! Protected routes (auth required)
	// Cars routes
	cars := api.Group("/cars", middleware.Protected())
	cars.Get("/", handlers.GetCars)
	cars.Get("/:id", handlers.GetCar)
	cars.Post("/create", handlers.CreateCar)
	cars.Put("/update/:id", handlers.UpdateCar)
	cars.Delete("/delete/:id", handlers.DeleteCar)

	// Customers routes
	customers := api.Group("/customers", middleware.Protected())
	customers.Get("/", handlers.GetCustomers)
	customers.Get("/:id", handlers.GetCustomer)
	customers.Post("/create", handlers.CreateCustomer)
	customers.Put("/update/:id", handlers.UpdateCustomer)
	customers.Delete("/delete/:id", handlers.DeleteCustomer)

	// Maintenance routes
	maintenance := api.Group("/maintenance", middleware.Protected())
	maintenance.Get("/", handlers.GetMaintenanceRecords)
	maintenance.Get("/:id", handlers.GetMaintenanceRecord)
	maintenance.Post("/create", handlers.CreateMaintenanceRecord)
	maintenance.Delete("/delete/:id", handlers.DeleteMaintenanceRecord)

	// Rentals routes
	rentals := api.Group("/rentals", middleware.Protected())
	rentals.Get("/", handlers.GetRentals)
	rentals.Get("/:id", handlers.GetRental)
	rentals.Post("/create", handlers.CreateRental)
	rentals.Put("/update/:id/status", handlers.UpdateRentalStatus)
}
