package main

import (
	"log"

	"tradexcore/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(configs.FiberConfig())
	app.Use(recover.New())
	app.Listen(":20220")
	log.Fatal("TradeX Core is serving")
}
