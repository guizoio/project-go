package check

import "github.com/gofiber/fiber/v2"

type CheckHandler struct{}

func NewCheckHandler() CheckHandler {
	return CheckHandler{}
}

func (ref CheckHandler) Check(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("success")
}
