package services

import "github.com/gofiber/fiber/v2"

type userServices interface {
	login(*fiber.Ctx) error
	register(*fiber.Ctx) error
}

func (s ServiceInstances) login(c *fiber.Ctx) error {
	return c.JSON("login")
}

func (s ServiceInstances) register(c *fiber.Ctx) error {
	return c.JSON("register")
}

func RegisterUserService(serviceInstances ServiceInstances) func(fiber.Router) {
	var s userServices = serviceInstances
	return func(router fiber.Router) {
		router.Post("/login", s.login)
		router.Post("/register", s.register)
	}
}
