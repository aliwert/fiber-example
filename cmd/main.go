package main

import (
	"log"
	"os"

	"github.com/aliwert/fiber-example/database"
	"github.com/gofiber/fiber/v2"
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

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
