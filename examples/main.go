package main

import (
	"log"

	"github.com/amanasmuei/toyyibpay.git/client"
	"github.com/amanasmuei/toyyibpay.git/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// Initialize ToyyibPay client
	toyyibClient := client.NewClient()

	// Route to create a new bill
	app.Post("/create-bill", handlers.CreateBillHandler(toyyibClient))

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
