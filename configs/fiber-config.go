package configs

import (
	"tradexcore/middlewares"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	}
}
