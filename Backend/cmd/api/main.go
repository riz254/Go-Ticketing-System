package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/riz254/Go-Ticketing-System.git/config"
	"github.com/riz254/Go-Ticketing-System.git/db"
	"github.com/riz254/Go-Ticketing-System.git/handlers"
	"github.com/riz254/Go-Ticketing-System.git/middlewares"
	"github.com/riz254/Go-Ticketing-System.git/repositories"
	"github.com/riz254/Go-Ticketing-System.git/services"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // You can replace * with specific origins for security
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorization",
	}))
	// Repositories
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	// Handlers
	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

	// Log and listen on port 3000
	fmt.Println("Listening on port 3000...")
	if err := app.Listen(fmt.Sprintf(":" + envConfig.ServerPort)); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
