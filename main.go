package main

import (
	"log"

	"tradexcore/configs"
	"tradexcore/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil { // Load .env file
		panic("Load .env file failed")
	}

	configs.ConnectRedisDB()                   // Connect Redis Database
	dbClient, dbContext := configs.ConnectDB() // Connect MongoDB Client
	defer dbClient.Disconnect(dbContext)       // Disconnect after end application

	app := fiber.New(configs.FiberConfig())
	app.Use(logger.New(logger.ConfigDefault)) // Logging Request

	app.Route("/", services.RegisterMainService) // Initial Router
	app.Use(recover.New())                       // Enable error handler inside fiber config

	app.Listen(":20220")                // Start web server
	log.Fatal("TradeX Core is serving") // Logging message
}
