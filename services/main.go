package services

import "github.com/gofiber/fiber/v2"

type Services interface {
	greeting(*fiber.Ctx) error
}

type services struct{}

func (s services) greeting(c *fiber.Ctx) error {
	return c.JSON("Hello Fiber")
}

func RegisterMainService(router fiber.Router) {
	var s Services = services{}
	router.Get("/", s.greeting)
}
