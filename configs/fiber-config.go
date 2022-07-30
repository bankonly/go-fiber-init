package configs

import (
	"gofiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	}
}
