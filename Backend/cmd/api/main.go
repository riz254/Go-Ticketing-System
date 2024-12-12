package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/riz254/Go-Ticketing-System.git/config"
	"github.com/riz254/Go-Ticketing-System.git/db"
	"github.com/riz254/Go-Ticketing-System.git/handlers"
	"github.com/riz254/Go-Ticketing-System.git/repositories"
)

func main() {
	fmt.Println("Starting the TicketBooking app...") // Log to check if the app starts

	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)
	if db == nil {
		fmt.Println("Failed to initialize database connection.")
		return
	}

	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)

	// Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	// Log and listen on port 3000
	fmt.Println("Listening on port 3000...")
	if err := app.Listen(fmt.Sprintf(":" + envConfig.ServerPort)); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
