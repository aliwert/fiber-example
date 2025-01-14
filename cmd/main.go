package main

import (
	"log"
	"os"

	"github.com/aliwert/fiber-example/database"
	"github.com/aliwert/fiber-example/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	// create Fiber app
	app := fiber.New()

	// Add middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, PATCH",
	}))

	// Setup routes
	routes.SetupRoutes(app)

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
